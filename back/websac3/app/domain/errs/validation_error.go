package errs

type validationError struct {
	baseError
}

func (e *validationError) Unwrap() error {
	return ValidationError
}

func NewValidationError(message string) error {
	return &validationError{baseError{msg: message}}
}
