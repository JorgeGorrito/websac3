package mapper

import (
	"errors"
	"fmt"
	"reflect"
)

func setSrcKeysValues(tagToField *map[string]reflect.Value, srcValue reflect.Value) {
	var srcType reflect.Type = srcValue.Type()
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
		srcType = srcValue.Type()
	}

	var field reflect.StructField
	var fieldTag string
	var srcValueField reflect.Value
	var isSrcValueFieldPtrToStruct bool
	var isSrcValuesAStruct bool
	for i := 0; i < srcType.NumField(); i++ {
		field = srcType.Field(i)
		fieldTag = field.Tag.Get("mapper")

		if fieldTag == "" {
			continue
		}

		srcValueField = srcValue.Field(i)
		isSrcValueFieldPtrToStruct = srcValueField.Kind() == reflect.Ptr && !srcValueField.IsNil() && srcValueField.Elem().Kind() == reflect.Struct
		isSrcValuesAStruct = srcValueField.Kind() == reflect.Struct
		if isSrcValueFieldPtrToStruct || isSrcValuesAStruct {
			setSrcKeysValues(tagToField, srcValueField)
		}
		(*tagToField)[fieldTag] = srcValueField
	}
}

func mapWithoutOverLoadTags(src, dest any, tagToField *map[string]reflect.Value) error {
	var srcValue reflect.Value = reflect.ValueOf(src)
	var destValue reflect.Value = reflect.ValueOf(dest)

	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be a pointer to a struct")
	}

	if srcValue.Kind() != reflect.Ptr || srcValue.Elem().Kind() != reflect.Struct {
		return errors.New("src must be a pointer to a struct")
	}

	var destElem reflect.Value = destValue.Elem()
	if tagToField == nil {
		tagToField = &map[string]reflect.Value{}
	}
	setSrcKeysValues(tagToField, srcValue)
	fmt.Printf("%+v\n", tagToField)

	var destType reflect.Type = destElem.Type()
	var field reflect.StructField
	var tag string
	var destField reflect.Value
	for i := 0; i < destElem.NumField(); i++ {
		field = destType.Field(i)
		tag = field.Tag.Get("mapper")
		if tag == "" {
			continue
		}

		srcField, exists := (*tagToField)[tag]
		if !exists || !field.IsExported() {
			continue
		}

		destField = destElem.Field(i)
		// Si los tipos son directamente asignables
		if srcField.Type().AssignableTo(destField.Type()) {
			destField.Set(srcField)
			continue
		}

		if srcField.Type().ConvertibleTo(destField.Type()) {
			destField.Set(srcField.Convert(destField.Type()))
			continue
		}

		// Si ambos son estructuras (o punteros a estructuras), intentamos mapearlas recursivamente
		srcFieldType := srcField.Type()
		destFieldType := destField.Type()

		// Manejar punteros a estructuras
		if srcFieldType.Kind() == reflect.Ptr && srcFieldType.Elem().Kind() == reflect.Struct &&
			destFieldType.Kind() == reflect.Ptr && destFieldType.Elem().Kind() == reflect.Struct {

			if srcField.IsNil() {
				continue // o manejar el caso nil como prefieras
			}

			if destField.IsNil() {
				destField.Set(reflect.New(destFieldType.Elem()))
			}

			continue
		}

		// Manejar estructuras directas (no punteros)
		if srcFieldType.Kind() == reflect.Struct && destFieldType.Kind() == reflect.Struct {
			// Creamos un nuevo valor para el destino
			newDestValue := reflect.New(destFieldType)
			if err := mapWithoutOverLoadTags(srcField.Addr().Interface(), newDestValue.Interface(), tagToField); err != nil {
				return err
			}
			destField.Set(newDestValue.Elem())
			continue
		}

		return fmt.Errorf(
			"incompatible types for tag '%s': %s vs %s",
			tag,
			srcField.Type(),
			destField.Type(),
		)
	}
	return nil
}

func Map(src, dest any) error {
	return mapWithoutOverLoadTags(src, dest, nil)
}
