package util

import (
	"errors"
	"websac3/app/domain/errs"
)

func GetResultMessageByErr(err error, defaultMessage string) string {
	switch {
	case errors.Is(err, errs.ValidationError) || errors.Is(err, errs.NotFoundError) || errors.Is(err, errs.ConflictError):
		return err.Error()
	}
	return defaultMessage
}
