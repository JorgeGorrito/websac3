package errs

import (
	"errors"
	"fmt"
)

var (
	ValidationError error = errors.New("validation error")
	ConflictError   error = errors.New("conflict error")
	NotFoundError   error = errors.New("not found error")
)

func NewValidationError(message string) error {
	return fmt.Errorf("%w: %s", ValidationError, message)
}

func NewConflictError(message string) error {
	return fmt.Errorf("%w: %s", ConflictError, message)
}

func NewNotFoundError(message string) error {
	return fmt.Errorf("%w: %s", NotFoundError, message)
}
