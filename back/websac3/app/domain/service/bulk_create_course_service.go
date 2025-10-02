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

type BulkCreateCourseService struct {
	persistenceManager _db.Manager
	createCoursePort   persistence.CreateCoursePort
	validator          validator.Validator
}

func NewBulkCreateCourseService(
	persistenceManager _db.Manager,
	createCoursePort persistence.CreateCoursePort,
	validator validator.Validator,
) *BulkCreateCourseService {
	return &BulkCreateCourseService{
		persistenceManager: persistenceManager,
		createCoursePort:   createCoursePort,
		validator:          validator,
	}
}

func (s *BulkCreateCourseService) Execute(
	bulkCommand command.BulkCreateCourseCommand,
	lang string,
) (response.BulkCreateCourseResponse, error) {
	result := response.BulkCreateCourseResponse{
		SuccessfulItems: []response.CourseBulkResult{},
		FailedItems:     []response.CourseBulkError{},
	}

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

	if len(records) <= 1 { // Skip header row
		return result, fmt.Errorf("CSV file is empty or contains only headers")
	}

	result.TotalProcessed = uint(len(records) - 1)

	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}
		rowNumber := uint(i + 1)

		// Parse and validate the row
		course, validationErrors := s.parseAndValidateRow(record, bulkCommand.DegreeProgramID, bulkCommand.CreatedBy, lang)
		if len(validationErrors) > 0 {
			var code *string
			var name *string
			if course.Code != "" {
				code = &course.Code
			}
			if course.Name != "" {
				name = &course.Name
			}
			result.FailedItems = append(result.FailedItems, response.CourseBulkError{
				RowNumber: rowNumber,
				Code:      code,
				Name:      name,
				Errors:    validationErrors,
			})
			result.FailedCount++
			continue
		}

		// Try to create the course
		err := s.persistenceManager.ExecuteInTransaction(
			func(ctx _db.Context) error {
				return s.createCoursePort.Create(&course, ctx)
			},
		)

		if err != nil {
			code := course.Code
			name := course.Name
			result.FailedItems = append(result.FailedItems, response.CourseBulkError{
				RowNumber: rowNumber,
				Code:      &code,
				Name:      &name,
				Errors:    []string{err.Error()},
			})
			result.FailedCount++
		} else {
			result.SuccessfulItems = append(result.SuccessfulItems, response.CourseBulkResult{
				RowNumber: rowNumber,
				Code:      course.Code,
				Name:      course.Name,
				Message:   "Course created successfully",
			})
			result.SuccessfulCount++
		}
	}

	return result, nil
}

func (s *BulkCreateCourseService) parseAndValidateRow(
	record []string,
	degreeProgramID uint,
	createdBy uint,
	lang string,
) (entity.Course, []string) {
	var course entity.Course
	var errors []string

	if len(record) < 7 {
		errors = append(errors, "Row does not contain enough columns (minimum 7 required)")
		return course, errors
	}

	// Name
	course.Name = strings.TrimSpace(record[0])
	if course.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Code
	course.Code = strings.TrimSpace(record[1])
	if course.Code == "" {
		errors = append(errors, "Code is required")
	}

	// Credits
	if creditsStr := strings.TrimSpace(record[2]); creditsStr != "" {
		if credits, err := strconv.ParseUint(creditsStr, 10, 32); err != nil {
			errors = append(errors, "Invalid credits format")
		} else {
			course.Credits = uint(credits)
		}
	} else {
		errors = append(errors, "Credits is required")
	}

	// Period Number
	if periodNumberStr := strings.TrimSpace(record[3]); periodNumberStr != "" {
		if periodNumber, err := strconv.ParseUint(periodNumberStr, 10, 32); err != nil {
			errors = append(errors, "Invalid period number format")
		} else {
			course.PeriodNumber = uint(periodNumber)
		}
	} else {
		errors = append(errors, "Period number is required")
	}

	// Nature ID
	if natureIDStr := strings.TrimSpace(record[4]); natureIDStr != "" {
		if natureID, err := strconv.ParseUint(natureIDStr, 10, 32); err != nil {
			errors = append(errors, "Invalid nature ID format")
		} else {
			course.NatureID = uint(natureID)
		}
	} else {
		errors = append(errors, "Nature ID is required")
	}

	// Type ID
	if typeIDStr := strings.TrimSpace(record[5]); typeIDStr != "" {
		if typeID, err := strconv.ParseUint(typeIDStr, 10, 32); err != nil {
			errors = append(errors, "Invalid type ID format")
		} else {
			course.TypeID = uint(typeID)
		}
	} else {
		errors = append(errors, "Type ID is required")
	}

	// Is Cybersecurity
	if isCybersecurityStr := strings.TrimSpace(strings.ToLower(record[6])); isCybersecurityStr != "" {
		if isCybersecurityStr == "true" || isCybersecurityStr == "1" || isCybersecurityStr == "yes" || isCybersecurityStr == "si" || isCybersecurityStr == "sí" {
			course.IsCybersecurity = true
		} else if isCybersecurityStr == "false" || isCybersecurityStr == "0" || isCybersecurityStr == "no" {
			course.IsCybersecurity = false
		} else {
			errors = append(errors, "Invalid is_cybersecurity format (use true/false, yes/no, si/no, 1/0)")
		}
	} else {
		course.IsCybersecurity = false
	}

	// Set additional fields
	course.DegreeProgramID = degreeProgramID
	course.CreatedBy = createdBy

	return course, errors
}

// ensureUTF8 converts content to UTF-8 if it's not already UTF-8
func (s *BulkCreateCourseService) ensureUTF8(content []byte) ([]byte, error) {
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
