package validator

import (
	"errors"
	"slices"
	"websac3/common/validator/errs"
)

func ValidateParamsRequired(params, paramsRequired []string) error {
	var errorList error
	for _, field := range paramsRequired {
		if !slices.Contains(params, field) {
			errorList = errors.Join(errorList, errs.NewFieldIsRequiredError(field))
		}
	}
	return errorList
}
