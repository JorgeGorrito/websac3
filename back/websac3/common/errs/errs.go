package errs

import "errors"

var ValidationError error = errors.New("validation error")
var ConflictError error = errors.New("conflict error")
