package errs

type conflictError struct {
	baseError
}

func (e *conflictError) Unwrap() error {
	return ConflictError
}

func NewConflictError(message string) error {
	return &conflictError{baseError{
		msg: message,
	}}
}
