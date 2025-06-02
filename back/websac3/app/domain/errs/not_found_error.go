package errs

type notFoundError struct {
	baseError
}

func (e *notFoundError) Unwrap() error {
	return NotFoundError
}

func NewNotFoundError(message string) error {
	return &notFoundError{baseError{msg: message}}
}
