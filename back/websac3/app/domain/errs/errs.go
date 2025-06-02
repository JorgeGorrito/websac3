package errs

import (
	"errors"
)

var (
	ValidationError error = errors.New("validation error")
	ConflictError   error = errors.New("conflict error")
	NotFoundError   error = errors.New("not found error")
)

type baseError struct {
	msg string
}

func (e *baseError) Error() string {
	return e.msg
}
