package errs

import (
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
)

type FieldIsRequiredError error

func NewFieldIsRequiredError(fieldName string, msgProvider message.Provider, lang string) FieldIsRequiredError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_is_required", fieldName),
	)
}

type FieldCantBeNullError error

func NewFieldCantBeNullError(fieldName string, msgProvider message.Provider, lang string) FieldCantBeNullError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_cant_be_null", fieldName),
	)
}

type ValueIsNotNumberError error

func NewValueIsNotNumberError(value string, msgProvider message.Provider, lang string) ValueIsNotNumberError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "value_is_not_number", value),
	)
}

type FieldValueIsNotNumberError error

func NewFieldValueIsNotNumberError(fieldName string, msgProvider message.Provider, lang string) FieldValueIsNotNumberError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_value_is_not_number", fieldName),
	)
}

type FieldMustBeGreaterThanError error

func NewFieldMustBeGreaterThanError(fieldName string, value float64, msgProvider message.Provider, lang string) FieldMustBeGreaterThanError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_must_be_greater_than", fieldName, value),
	)
}

type FieldMustBeLessThanError error

func NewFieldMustBeLessThanError(fieldName string, value float64, msgProvider message.Provider, lang string) FieldMustBeLessThanError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_must_be_less_than", fieldName, value),
	)
}

type FieldMustBeDifferentToError error

func NewFieldMustBeDifferentToError(fieldName string, value string, msgProvider message.Provider, lang string) FieldMustBeDifferentToError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_must_be_different_to", fieldName, value),
	)
}

type FieldMustBeEmailError error

func NewFieldMustBeEmailError(fieldName string, msgProvider message.Provider, lang string) FieldMustBeEmailError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_must_be_email", fieldName),
	)
}

type FieldMinLengthError error

func NewFieldMinLengthError(fieldName string, minLength int, msgProvider message.Provider, lang string) FieldMinLengthError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_min_length", fieldName, minLength),
	)
}

type FieldMaxLengthError error

func NewFieldMaxLengthError(fieldName string, maxLength int, msgProvider message.Provider,
	lang string) FieldMaxLengthError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_max_length", fieldName, maxLength),
	)
}

type FieldMustBeStringError error

func NewFieldMustBeStringError(fieldName string, msgProvider message.Provider, lang string) FieldMustBeStringError {
	return errs.NewValidationError(
		msgProvider.
			WithLang(lang).
			GetMessage("validator", "field_must_be_string", fieldName),
	)
}
