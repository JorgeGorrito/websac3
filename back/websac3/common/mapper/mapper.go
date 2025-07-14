package mapper

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	SrcPointerIsNilError = errors.New("source pointer is nil")
)

type MapFunc func(sourcePtr any) (any, error)

var registry map[reflect.Type]map[reflect.Type]MapFunc = make(map[reflect.Type]map[reflect.Type]MapFunc)

func BuildMapFunc[S any, D any](fn func(*S) (D, error)) MapFunc {
	return func(sourcePtr any) (any, error) {
		typedSrc, ok := sourcePtr.(*S)
		if !ok {
			return nil, fmt.Errorf("invalid source type: expected *%T", new(S))
		}
		return fn(typedSrc)
	}
}

func RegisterMapFunc[S any, D any](fn func(*S) (D, error)) {
	var zero D
	srcType := reflect.TypeOf((*S)(nil)).Elem()
	dstType := reflect.TypeOf(zero)

	if _, ok := registry[srcType]; !ok {
		registry[srcType] = make(map[reflect.Type]MapFunc)
	}
	registry[srcType][dstType] = BuildMapFunc(fn)
}

func Map[S any, D any](sourcePtr *S) (D, error) {
	var zero D

	if sourcePtr == nil {
		return zero, SrcPointerIsNilError
	}

	srcType := reflect.TypeOf((*S)(nil)).Elem()
	var d D
	dstType := reflect.TypeOf(d)

	mapFuncs, found := registry[srcType]
	if !found {
		return zero, fmt.Errorf("no mapping functions registered for type %v", srcType)
	}

	mapFunc, found := mapFuncs[dstType]
	if !found {
		return zero, fmt.Errorf("no mapping function from %v to %v", srcType, dstType)
	}

	result, err := mapFunc(sourcePtr)
	if err != nil {
		return zero, err
	}

	typedResult, ok := result.(D)
	if !ok {
		return zero, fmt.Errorf("invalid result type")
	}

	return typedResult, nil
}
