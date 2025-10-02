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

type BulkCreateCourseTopicService struct {
	persistenceManager    _db.Manager
	createCourseTopicPort persistence.CreateCourseTopicPort
	getCoursePort         persistence.GetCoursePort
	getUserPort           persistence.GetUserPort
	validator             validator.Validator
}

func NewBulkCreateCourseTopicService(
	persistenceManager _db.Manager,
	createCourseTopicPort persistence.CreateCourseTopicPort,
	getCoursePort persistence.GetCoursePort,
	getUserPort persistence.GetUserPort,
	validator validator.Validator,
) *BulkCreateCourseTopicService {
	return &BulkCreateCourseTopicService{
		persistenceManager:    persistenceManager,
		createCourseTopicPort: createCourseTopicPort,
		getCoursePort:         getCoursePort,
		getUserPort:           getUserPort,
		validator:             validator,
	}
}

func (s *BulkCreateCourseTopicService) Execute(
	bulkCommand command.BulkCreateCourseTopicCommand,
	lang string,
) (response.BulkCreateCourseTopicResponse, error) {
	result := response.BulkCreateCourseTopicResponse{
		SuccessfulItems: []response.CourseTopicBulkResult{},
		FailedItems:     []response.CourseTopicBulkError{},
	}

	// Verify that the course exists and that the user has permission to modify it
	var course *entity.Course
	var user entity.User
	err := s.persistenceManager.ExecuteInTransaction(
		func(ctx _db.Context) error {
			var err error
			course, err = s.getCoursePort.GetByID(bulkCommand.CourseID, ctx)
			if err != nil {
				return fmt.Errorf("course not found or cannot be accessed: %w", err)
			}

			// Get the user to check permissions
			user, err = s.getUserPort.GetByID(bulkCommand.UserID, ctx)
			if err != nil {
				return fmt.Errorf("user not found: %w", err)
			}

			// Check if user has permission to modify this course
			if !user.IsAdmin() && !course.CanBeUpdatedBy(bulkCommand.UserID) {
				return fmt.Errorf("user does not have permission to modify this course")
			}

			return nil
		},
	)
	if err != nil {
		return result, err
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
		courseTopic, validationErrors := s.parseAndValidateRow(record, bulkCommand.CourseID, lang)
		if len(validationErrors) > 0 {
			var topicID *uint
			var studyHours *uint
			if courseTopic.TopicID != 0 {
				topicID = &courseTopic.TopicID
			}
			if courseTopic.StudyHours != 0 {
				sh := uint(courseTopic.StudyHours)
				studyHours = &sh
			}
			result.FailedItems = append(result.FailedItems, response.CourseTopicBulkError{
				RowNumber:  rowNumber,
				TopicID:    topicID,
				StudyHours: studyHours,
				Errors:     validationErrors,
			})
			result.FailedCount++
			continue
		}

		// Try to create the course topic
		err := s.persistenceManager.ExecuteInTransaction(
			func(ctx _db.Context) error {
				return s.createCourseTopicPort.CreateCourseTopic(&courseTopic, ctx)
			},
		)

		if err != nil {
			tid := courseTopic.TopicID
			sh := uint(courseTopic.StudyHours)
			result.FailedItems = append(result.FailedItems, response.CourseTopicBulkError{
				RowNumber:  rowNumber,
				TopicID:    &tid,
				StudyHours: &sh,
				Errors:     []string{err.Error()},
			})
			result.FailedCount++
		} else {
			result.SuccessfulItems = append(result.SuccessfulItems, response.CourseTopicBulkResult{
				RowNumber:  rowNumber,
				TopicID:    courseTopic.TopicID,
				StudyHours: uint(courseTopic.StudyHours),
				Message:    "Course topic added successfully",
			})
			result.SuccessfulCount++
		}
	}

	return result, nil
}

func (s *BulkCreateCourseTopicService) parseAndValidateRow(
	record []string,
	courseID uint,
	lang string,
) (entity.CourseTopic, []string) {
	var courseTopic entity.CourseTopic
	var errors []string

	if len(record) < 2 {
		errors = append(errors, "Row does not contain enough columns (minimum 2 required)")
		return courseTopic, errors
	}

	// Topic ID
	if topicIDStr := strings.TrimSpace(record[0]); topicIDStr != "" {
		if topicID, err := strconv.ParseUint(topicIDStr, 10, 32); err != nil {
			errors = append(errors, "Invalid topic ID format")
		} else {
			courseTopic.TopicID = uint(topicID)
		}
	} else {
		errors = append(errors, "Topic ID is required")
	}

	// Study Hours
	if studyHoursStr := strings.TrimSpace(record[1]); studyHoursStr != "" {
		if studyHours, err := strconv.ParseFloat(studyHoursStr, 32); err != nil {
			errors = append(errors, "Invalid study hours format")
		} else {
			courseTopic.StudyHours = float32(studyHours)
		}
	} else {
		errors = append(errors, "Study hours is required")
	}

	// Set course ID
	courseTopic.CourseID = courseID

	return courseTopic, errors
}

// ensureUTF8 converts content to UTF-8 if it's not already UTF-8
func (s *BulkCreateCourseTopicService) ensureUTF8(content []byte) ([]byte, error) {
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
