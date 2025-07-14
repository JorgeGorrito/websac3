package errs

import (
	"fmt"
	"websac3/app/domain/errs"
)

type FieldIsRequiredError error

func NewFieldIsRequiredError(fieldName string) FieldIsRequiredError {
	return errs.NewValidationError(fmt.Sprintf("field %s is required.", fieldName))
}

type FieldCantBeNullError error

func NewFieldCantBeNullError(fieldName string) FieldCantBeNullError {
	return errs.NewValidationError(fmt.Sprintf("field %s can't be null.", fieldName))
}

type ValueIsNotNumberError error

func NewValueIsNotNumberError(value string) ValueIsNotNumberError {
	return errs.NewValidationError(fmt.Sprintf("value '%s' is not a number.", value))
}

type FieldValueIsNotNumberError error

func NewFieldValueIsNotNumberError(fieldName string) FieldValueIsNotNumberError {
	return errs.NewValidationError(fmt.Sprintf("field %s value is not a number.", fieldName))
}

type FieldMustBeGreaterThanError error

func NewFieldMustBeGreaterThanError(fieldName string, value float64) FieldMustBeGreaterThanError {
	return errs.NewValidationError(fmt.Sprintf("field %s must be greater than %f.", fieldName, value))
}

type FieldMustBeLessThanError error

func NewFieldMustBeLessThanError(fieldName string, value float64) FieldMustBeLessThanError {
	return errs.NewValidationError(fmt.Sprintf("field %s must be less than %f.", fieldName, value))
}

type FieldMustBeDifferentToError error

func NewFieldMustBeDifferentToError(fieldName string, value string) FieldMustBeDifferentToError {
	return errs.NewValidationError(fmt.Sprintf("field %s must be different to '%s'.", fieldName, value))
}

type FieldMustBeEmailError error

func NewFieldMustBeEmailError(fieldName string) FieldMustBeEmailError {
	return errs.NewValidationError(fmt.Sprintf("field %s must be a valid email.", fieldName))
}
