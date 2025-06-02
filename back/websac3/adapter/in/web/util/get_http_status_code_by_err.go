package util

import (
	"errors"
	"net/http"
	"websac3/app/domain/errs"
)

func GetHttpStatusCodeByErr(err error) (httpStatusCode int) {
	switch {
	case errors.Is(err, errs.ValidationError):
		return http.StatusBadRequest
	case errors.Is(err, errs.NotFoundError):
		return http.StatusNotFound
	case errors.Is(err, errs.ConflictError):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
