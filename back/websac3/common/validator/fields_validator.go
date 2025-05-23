package validator

import (
	"errors"
	"fmt"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
	verrs "websac3/common/validator/errs"
)

type FieldValidationFunction func(string, reflect.Value, string) error

func validateRequired(fieldName string, objectValue reflect.Value, _ string) error {
	if objectValue.IsZero() {
		return verrs.NewFieldIsRequiredError(fieldName)
	}
	return nil
}

func validateNotNull(fieldName string, objectValue reflect.Value, _ string) error {
	if objectValue.Kind() == reflect.Ptr && objectValue.IsNil() {
		return verrs.NewFieldCantBeNullError(fieldName)
	}
	return nil
}

func validateGreaterThan(fieldName string, objectValue reflect.Value, value string) error {
	valueAsFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return verrs.NewValueIsNotNumberError(value)
	}
	objectValueAsString := fmt.Sprintf("%v", objectValue)
	objectValueAsFloat, err := strconv.ParseFloat(objectValueAsString, 64)
	if err != nil {
		return verrs.NewFieldValueIsNotNumberError(fieldName)
	}
	if objectValueAsFloat <= valueAsFloat {
		return verrs.NewFieldMustBeGreaterThanError(fieldName, objectValueAsFloat)
	}
	return nil
}

func validateLessThan(fieldName string, objectValue reflect.Value, value string) error {
	valueAsFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return verrs.NewValueIsNotNumberError(value)
	}
	objectValueAsString := fmt.Sprintf("%v", objectValue)
	objectValueAsFloat, err := strconv.ParseFloat(objectValueAsString, 64)
	if err != nil {
		return verrs.NewFieldValueIsNotNumberError(fieldName)
	}
	if objectValueAsFloat >= valueAsFloat {
		return verrs.NewFieldMustBeLessThanError(fieldName, objectValueAsFloat)
	}
	return nil
}

func validateDiffTo(fieldName string, objectValue reflect.Value, value string) error {
	if objectValue.String() == value {
		return verrs.NewFieldMustBeDifferentToError(fieldName, value)
	}
	return nil
}

func validateIsEmail(fieldName string, objectValue reflect.Value, value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return verrs.NewFieldMustBeEmailError(fieldName)
	}

	return nil
}

var fieldsValidationRules = map[string]FieldValidationFunction{
	"required":     validateRequired,
	"not_null":     validateNotNull,
	"greater_than": validateGreaterThan,
	"less_than":    validateLessThan,
	"diff_to":      validateDiffTo,
	"email":        validateIsEmail,
}

func ValidateFields(t interface{}) error {
	var errs error
	var val reflect.Value = reflect.ValueOf(t).Elem()
	var typ reflect.Type = val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		validationsTag := field.Tag.Get("validations")
		if validationsTag == "" {
			continue
		}

		var validationsWithValues [][2]string
		for _, validationRule := range strings.Split(validationsTag, ";") {
			switch validation := strings.Split(validationRule, "="); {
			case len(validation) == 2:
				validationsWithValues = append(validationsWithValues, [2]string{validation[0], validation[1]})
			case len(validation) == 1:
				validationsWithValues = append(validationsWithValues, [2]string{validation[0], ""})
			default:
				continue
			}
		}

		fieldValue := val.Field(i)
		const (
			validationName = iota
			validationValue
		)
		for _, validation := range validationsWithValues {
			fieldValidationFunc, ok := fieldsValidationRules[validation[validationName]]
			if !ok {
				continue
			}
			if err := fieldValidationFunc(field.Name, fieldValue, validation[validationValue]); err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}
	return errs
}
