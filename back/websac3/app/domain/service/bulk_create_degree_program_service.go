package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/out/persistence"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/validator"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type BulkCreateDegreeProgramService struct {
	persistenceManager      _db.Manager
	createDegreeProgramPort persistence.CreateDegreeProgramPort
	getProfessionalRolePort persistence.GetProfessionalRolePort
	getDurationUnitPort     persistence.GetDurationUnitPort
	getFormationLevelPort   persistence.GetFormationLevelPort
	validator               validator.Validator
}

func NewBulkCreateDegreeProgramService(
	persistenceManager _db.Manager,
	createDegreeProgramPort persistence.CreateDegreeProgramPort,
	getProfessionalRolePort persistence.GetProfessionalRolePort,
	getDurationUnitPort persistence.GetDurationUnitPort,
	getFormationLevelPort persistence.GetFormationLevelPort,
	validator validator.Validator,
) *BulkCreateDegreeProgramService {
	return &BulkCreateDegreeProgramService{
		persistenceManager:      persistenceManager,
		createDegreeProgramPort: createDegreeProgramPort,
		getProfessionalRolePort: getProfessionalRolePort,
		getDurationUnitPort:     getDurationUnitPort,
		getFormationLevelPort:   getFormationLevelPort,
		validator:               validator,
	}
}

func (s *BulkCreateDegreeProgramService) Execute(
	bulkCommand command.BulkCreateDegreeProgramCommand,
	lang string,
) (response.BulkCreateDegreeProgramResponse, error) {
	var result response.BulkCreateDegreeProgramResponse
	result.SuccessfulItems = []response.DegreeProgramBulkResult{}
	result.FailedItems = []response.DegreeProgramBulkError{}

	// Open the uploaded file
	file, err := bulkCommand.File.Open()
	if err != nil {
		return result, fmt.Errorf("error opening uploaded file: %w", err)
	}
	defer file.Close()

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		return result, fmt.Errorf("error reading uploaded file: %w", err)
	}

	// Convert to UTF-8 if needed
	utf8Content, err := s.ensureUTF8(content)
	if err != nil {
		return result, fmt.Errorf("error converting file encoding: %w", err)
	}

	// Parse CSV with semicolon separator for Excel compatibility
	reader := csv.NewReader(bytes.NewReader(utf8Content))
	reader.Comma = ';'          // Use semicolon as separator
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return result, fmt.Errorf("error reading CSV file: %w", err)
	}

	if len(records) < 2 {
		return result, fmt.Errorf("CSV file must contain at least a header row and one data row")
	}

	// Skip header row
	dataRows := records[1:]
	result.TotalProcessed = len(dataRows)

	// Process each row
	for rowIndex, record := range dataRows {
		rowNumber := rowIndex + 2 // +2 because we skip header and rowIndex is 0-based

		if len(record) < 11 {
			result.FailedItems = append(result.FailedItems, response.DegreeProgramBulkError{
				RowNumber: rowNumber,
				Errors:    []string{"Row does not contain enough columns (minimum 11 required)"},
			})
			result.FailedCount++
			continue
		}

		// Parse and validate the row
		degreeProgram, validationErrors := s.parseAndValidateRow(record, bulkCommand.CreatedBy)
		if len(validationErrors) > 0 {
			var snies *uint
			var name *string
			if degreeProgram.Snies > 0 {
				snies = &degreeProgram.Snies
			}
			if degreeProgram.Name != "" {
				name = &degreeProgram.Name
			}
			result.FailedItems = append(result.FailedItems, response.DegreeProgramBulkError{
				RowNumber: rowNumber,
				Snies:     snies,
				Name:      name,
				Errors:    validationErrors,
			})
			result.FailedCount++
			continue
		}

		// Try to create the degree program
		err := s.persistenceManager.ExecuteInTransaction(
			func(ctx _db.Context) error {
				// Validate that duration unit ID exists
				if degreeProgram.DurationUnitID > 0 {
					if _, err := s.getDurationUnitPort.GetByID(degreeProgram.DurationUnitID, ctx); err != nil {
						return fmt.Errorf("duration unit with ID %d not found", degreeProgram.DurationUnitID)
					}
				}

				// Validate that formation level ID exists
				if degreeProgram.FormationLevelID > 0 {
					if _, err := s.getFormationLevelPort.GetByID(degreeProgram.FormationLevelID, ctx); err != nil {
						return fmt.Errorf("formation level with ID %d not found", degreeProgram.FormationLevelID)
					}
				}

				// Validate that all professional role IDs exist
				if len(degreeProgram.ProfessionalRoles) > 0 {
					var validatedRoles []entity.ProfessionalRole
					for _, pr := range degreeProgram.ProfessionalRoles {
						role, err := s.getProfessionalRolePort.GetByID(pr.ID, lang, ctx)
						if err != nil {
							return fmt.Errorf("professional role with ID %d not found", pr.ID)
						}
						validatedRoles = append(validatedRoles, role)
					}
					degreeProgram.ProfessionalRoles = validatedRoles
				}

				return s.createDegreeProgramPort.Create(&degreeProgram, ctx)
			},
		)

		if err != nil {
			snies := degreeProgram.Snies
			name := degreeProgram.Name
			result.FailedItems = append(result.FailedItems, response.DegreeProgramBulkError{
				RowNumber: rowNumber,
				Snies:     &snies,
				Name:      &name,
				Errors:    []string{err.Error()},
			})
			result.FailedCount++
		} else {
			result.SuccessfulItems = append(result.SuccessfulItems, response.DegreeProgramBulkResult{
				RowNumber: rowNumber,
				Snies:     degreeProgram.Snies,
				Name:      degreeProgram.Name,
				Message:   "Degree program created successfully",
			})
			result.SuccessfulCount++
		}
	}

	return result, nil
}

func (s *BulkCreateDegreeProgramService) parseAndValidateRow(
	record []string,
	createdBy uint,
) (entity.DegreeProgram, []string) {
	var degreeProgram entity.DegreeProgram
	var errors []string

	// Parse SNIES
	if sniesStr := strings.TrimSpace(record[0]); sniesStr != "" {
		if snies, err := strconv.ParseUint(sniesStr, 10, 32); err != nil {
			errors = append(errors, "Invalid SNIES format")
		} else {
			degreeProgram.Snies = uint(snies)
		}
	} else {
		errors = append(errors, "SNIES is required")
	}

	// Parse Name
	degreeProgram.Name = strings.TrimSpace(record[1])
	if degreeProgram.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Parse Total Credits
	if creditsStr := strings.TrimSpace(record[2]); creditsStr != "" {
		if credits, err := strconv.ParseUint(creditsStr, 10, 32); err != nil {
			errors = append(errors, "Invalid total credits format")
		} else {
			degreeProgram.TotalCredits = uint(credits)
		}
	} else {
		errors = append(errors, "Total credits is required")
	}

	// Parse Duration Value
	if durationValueStr := strings.TrimSpace(record[3]); durationValueStr != "" {
		if durationValue, err := strconv.ParseUint(durationValueStr, 10, 32); err != nil {
			errors = append(errors, "Invalid duration value format")
		} else {
			degreeProgram.DurationValue = uint(durationValue)
		}
	} else {
		errors = append(errors, "Duration value is required")
	}

	// Parse Duration Unit ID
	if durationUnitIDStr := strings.TrimSpace(record[4]); durationUnitIDStr != "" {
		if durationUnitID, err := strconv.ParseUint(durationUnitIDStr, 10, 32); err != nil {
			errors = append(errors, "Invalid duration unit ID format")
		} else {
			degreeProgram.DurationUnitID = uint(durationUnitID)
		}
	} else {
		errors = append(errors, "Duration unit ID is required")
	}

	// Parse Formation Level ID
	if formationLevelIDStr := strings.TrimSpace(record[5]); formationLevelIDStr != "" {
		if formationLevelID, err := strconv.ParseUint(formationLevelIDStr, 10, 32); err != nil {
			errors = append(errors, "Invalid formation level ID format")
		} else {
			degreeProgram.FormationLevelID = uint(formationLevelID)
		}
	} else {
		errors = append(errors, "Formation level ID is required")
	}

	// Parse Professional Role IDs
	if professionalRoleIDsStr := strings.TrimSpace(record[6]); professionalRoleIDsStr != "" {
		roleIDStrings := strings.Split(professionalRoleIDsStr, ",")
		var professionalRoleIDs []uint
		for _, roleIDStr := range roleIDStrings {
			if roleID, err := strconv.ParseUint(strings.TrimSpace(roleIDStr), 10, 32); err != nil {
				errors = append(errors, fmt.Sprintf("Invalid professional role ID format: %s", roleIDStr))
			} else {
				professionalRoleIDs = append(professionalRoleIDs, uint(roleID))
			}
		}
		if len(professionalRoleIDs) == 0 {
			errors = append(errors, "At least one professional role ID is required")
		} else {
			// Create professional roles entities with the provided IDs
			// Database foreign key constraints will validate that these IDs exist
			degreeProgram.ProfessionalRoles = make([]entity.ProfessionalRole, len(professionalRoleIDs))
			for i, roleID := range professionalRoleIDs {
				degreeProgram.ProfessionalRoles[i] = entity.ProfessionalRole{ID: roleID}
			}
		}
	} else {
		errors = append(errors, "Professional role IDs are required")
	}

	// Parse Program Focus
	degreeProgram.ProgramFocus = strings.TrimSpace(record[7])
	if degreeProgram.ProgramFocus == "" {
		errors = append(errors, "Program focus is required")
	}

	// Parse Entry Profile
	degreeProgram.EntryProfile = strings.TrimSpace(record[8])
	if degreeProgram.EntryProfile == "" {
		errors = append(errors, "Entry profile is required")
	}

	// Parse Graduate Profile
	degreeProgram.GraduateProfile = strings.TrimSpace(record[9])
	if degreeProgram.GraduateProfile == "" {
		errors = append(errors, "Graduate profile is required")
	}

	// Parse Professional Profile
	degreeProgram.ProfessionalProfile = strings.TrimSpace(record[10])
	if degreeProgram.ProfessionalProfile == "" {
		errors = append(errors, "Professional profile is required")
	}

	// Set CreatedBy
	degreeProgram.CreatedBy = createdBy

	return degreeProgram, errors
}

// ensureUTF8 converts content to UTF-8 if it's not already UTF-8
func (s *BulkCreateDegreeProgramService) ensureUTF8(content []byte) ([]byte, error) {
	// Check if content is already valid UTF-8
	if utf8.Valid(content) {
		return content, nil
	}

	// Try to decode from Windows-1252 (common Excel encoding)
	decoder := charmap.Windows1252.NewDecoder()
	utf8Content, _, err := transform.Bytes(decoder, content)
	if err != nil {
		// If Windows-1252 fails, try ISO-8859-1 (Latin-1)
		decoder = charmap.ISO8859_1.NewDecoder()
		utf8Content, _, err = transform.Bytes(decoder, content)
		if err != nil {
			return nil, fmt.Errorf("unable to decode file content to UTF-8: %w", err)
		}
	}

	return utf8Content, nil
}
