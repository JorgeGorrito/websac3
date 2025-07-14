package validator

import (
	"errors"
	"fmt"
	"net/mail"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"websac3/app/port/out/message"
	"websac3/common/validator/errs"
	verrs "websac3/common/validator/errs"
)

type Validator interface {
	ValidateFields(t any, lang string) []error
	ValidateParamsRequired(params, paramsRequired []string, lang string) error
}

type FieldValidationFunction func(string, string, reflect.Value, string) error

type validator struct {
	msgProvider           message.Provider
	fieldsValidationRules map[string]FieldValidationFunction
}

func (v *validator) validateRequired(lang string, fieldName string, objectValue reflect.Value, _ string) error {
	if objectValue.IsZero() {
		return verrs.NewFieldIsRequiredError(fieldName, v.msgProvider, lang)
	}
	return nil
}

func (v *validator) validateNotNull(lang string, fieldName string, objectValue reflect.Value, _ string) error {
	if objectValue.Kind() == reflect.Ptr && objectValue.IsNil() {
		return verrs.NewFieldCantBeNullError(fieldName, v.msgProvider, lang)
	}
	return nil
}

func (v *validator) validateGreaterThan(lang string, fieldName string, objectValue reflect.Value, value string) error {
	valueAsFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return verrs.NewValueIsNotNumberError(value, v.msgProvider, lang)
	}
	objectValueAsString := fmt.Sprintf("%v", objectValue)
	objectValueAsFloat, err := strconv.ParseFloat(objectValueAsString, 64)
	if err != nil {
		return verrs.NewFieldValueIsNotNumberError(fieldName, v.msgProvider, lang)
	}
	if objectValueAsFloat <= valueAsFloat {
		return verrs.NewFieldMustBeGreaterThanError(fieldName, objectValueAsFloat, v.msgProvider, lang)
	}
	return nil
}

func (v *validator) validateLessThan(lang string, fieldName string, objectValue reflect.Value, value string) error {
	valueAsFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return verrs.NewValueIsNotNumberError(value, v.msgProvider, lang)
	}
	objectValueAsString := fmt.Sprintf("%v", objectValue)
	objectValueAsFloat, err := strconv.ParseFloat(objectValueAsString, 64)
	if err != nil {
		return verrs.NewFieldValueIsNotNumberError(fieldName, v.msgProvider, lang)
	}
	if objectValueAsFloat >= valueAsFloat {
		return verrs.NewFieldMustBeLessThanError(fieldName, objectValueAsFloat, v.msgProvider, lang)
	}
	return nil
}

func (v *validator) validateDiffTo(lang string, fieldName string, objectValue reflect.Value, value string) error {
	if objectValue.String() == value {
		return verrs.NewFieldMustBeDifferentToError(fieldName, value, v.msgProvider, lang)
	}
	return nil
}

func (v *validator) validateIsEmail(lang string, fieldName string, objectValue reflect.Value, _ string) error {
	_, err := mail.ParseAddress(fmt.Sprintf("%v", objectValue.Interface()))
	if err != nil {
		return verrs.NewFieldMustBeEmailError(fieldName, v.msgProvider, lang)
	}
	return nil
}

func (v *validator) ValidateParamsRequired(params, paramsRequired []string, lang string) error {
	var errorList error
	for _, field := range paramsRequired {
		if !slices.Contains(params, field) {
			errorList = errors.Join(errorList, errs.NewFieldIsRequiredError(field, v.msgProvider, lang))
		}
	}
	return errorList
}

func (v *validator) ValidateFields(t interface{}, lang string) []error {
	var errs []error
	var val reflect.Value = reflect.ValueOf(t).Elem()
	var typ reflect.Type = val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		if field.Type.Kind() == reflect.Struct {
			errsFound := v.ValidateFields(fieldValue.Addr().Interface(), lang)
			if len(errsFound) > 0 {
				errs = append(errs, errsFound...)
			}
		}

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

		const (
			validationName = iota
			validationValue
		)
		for _, validation := range validationsWithValues {
			fieldValidationFunc, ok := v.fieldsValidationRules[validation[validationName]]
			if !ok {
				continue
			}
			if err := fieldValidationFunc(lang, field.Name, fieldValue, validation[validationValue]); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}

func New(msgProvider message.Provider) *validator {
	v := &validator{
		msgProvider: msgProvider,
	}

	fieldValidationRules := map[string]FieldValidationFunction{
		"required":     v.validateRequired,
		"not_null":     v.validateNotNull,
		"greater_than": v.validateGreaterThan,
		"less_than":    v.validateLessThan,
		"diff_to":      v.validateDiffTo,
		"email":        v.validateIsEmail,
	}
	v.fieldsValidationRules = fieldValidationRules

	return v
}
