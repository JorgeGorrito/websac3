package errs

import (
	"fmt"
	"websac3/app/domain/errs"
)

type FieldIsRequiredError error

func NewFieldIsRequiredError(fieldName string) FieldIsRequiredError {
	return fmt.Errorf("field %s is required: %w", fieldName, errs.ValidationError)
}

type FieldCantBeNullError error

func NewFieldCantBeNullError(fieldName string) FieldCantBeNullError {
	return fmt.Errorf("field %s cannot be null: %w", fieldName, errs.ValidationError)
}

type ValueIsNotNumberError error

func NewValueIsNotNumberError(value string) ValueIsNotNumberError {
	return fmt.Errorf("value '%s' is not a number: %w", value, errs.ValidationError)
}

type FieldValueIsNotNumberError error

func NewFieldValueIsNotNumberError(fieldName string) FieldValueIsNotNumberError {
	return fmt.Errorf("field %s value is not a number: %w", fieldName, errs.ValidationError)
}

type FieldMustBeGreaterThanError error

func NewFieldMustBeGreaterThanError(fieldName string, value float64) FieldMustBeGreaterThanError {
	return fmt.Errorf("field %s must be greater than %f: %w", fieldName, value, errs.ValidationError)
}

type FieldMustBeLessThanError error

func NewFieldMustBeLessThanError(fieldName string, value float64) FieldMustBeLessThanError {
	return fmt.Errorf("field %s must be less than %f:%w", fieldName, value, errs.ValidationError)
}

type FieldMustBeDifferentToError error

func NewFieldMustBeDifferentToError(fieldName string, value string) FieldMustBeDifferentToError {
	return fmt.Errorf("field %s must be different to '%s':%w", fieldName, value, errs.ValidationError)
}

type FieldMustBeEmailError error

func NewFieldMustBeEmailError(fieldName string) FieldMustBeEmailError {
	return fmt.Errorf("field %s must be a valid email address: %w", fieldName, errs.ValidationError)
}
