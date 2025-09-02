package repository

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// IsUniqueConstraintViolation checks if the error is a PostgreSQL unique constraint violation
func IsUniqueConstraintViolation(err error) bool {
	if err == nil {
		return false
	}

	// Check for PostgreSQL error code 23505 (unique_violation)
	return strings.Contains(err.Error(), "23505") ||
		strings.Contains(err.Error(), "duplicate key value") ||
		strings.Contains(err.Error(), "unique constraint") ||
		strings.Contains(err.Error(), "already exists")
}

// IsForeignKeyViolation checks if the error is a PostgreSQL foreign key violation
func IsForeignKeyViolation(err error) bool {
	if err == nil {
		return false
	}

	// Check for PostgreSQL error code 23503 (foreign_key_violation)
	return strings.Contains(err.Error(), "23503") ||
		strings.Contains(err.Error(), "foreign key") ||
		strings.Contains(err.Error(), "referenced")
}

// IsCheckViolation checks if the error is a PostgreSQL check constraint violation
func IsCheckViolation(err error) bool {
	if err == nil {
		return false
	}

	// Check for PostgreSQL error code 23514 (check_violation)
	return strings.Contains(err.Error(), "23514") ||
		strings.Contains(err.Error(), "check constraint")
}

// IsNotNullViolation checks if the error is a PostgreSQL not null violation
func IsNotNullViolation(err error) bool {
	if err == nil {
		return false
	}

	// Check for PostgreSQL error code 23502 (not_null_violation)
	return strings.Contains(err.Error(), "23502") ||
		strings.Contains(err.Error(), "null value in column")
}

// IsRecordNotFound checks if the error is GORM's record not found
func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// GetPostgreSQLErrorCode extracts the PostgreSQL error code from the error message
func GetPostgreSQLErrorCode(err error) string {
	if err == nil {
		return ""
	}

	// PostgreSQL error codes are typically in format: ERROR: (23505)
	// or: pq: duplicate key value violates unique constraint
	errorStr := err.Error()

	// Look for error code pattern
	if strings.Contains(errorStr, "23505") {
		return "23505" // unique_violation
	}
	if strings.Contains(errorStr, "23503") {
		return "23503" // foreign_key_violation
	}
	if strings.Contains(errorStr, "23514") {
		return "23514" // check_violation
	}
	if strings.Contains(errorStr, "23502") {
		return "23502" // not_null_violation
	}

	return ""
}

// IsConflictError checks if the error represents any type of conflict
func IsConflictError(err error) bool {
	return IsUniqueConstraintViolation(err) ||
		IsForeignKeyViolation(err) ||
		IsCheckViolation(err)
}
