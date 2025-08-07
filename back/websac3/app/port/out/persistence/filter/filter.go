package filter

import (
	"slices"
	"websac3/app/port/out/persistence/db"
)

type Operator string

const (
	EqualOperator    Operator = "eq"
	ContainsOperator Operator = "cont"
)

type Filter struct {
	Field    string
	Operator Operator
	Value    any
}

type FilterFunc func(ctx db.Context, field string, value any) (db.Context, error)
type FilterFuncMap map[Operator]FilterFunc
type Filters []Filter

func (f *Filters) Apply(baseCtx db.Context, registryFilters FilterFuncMap) (db.Context, error) {
	var err error
	ctx := baseCtx
	for _, filter := range *f {
		if fn, ok := registryFilters[filter.Operator]; ok {
			ctx, err = fn(ctx, filter.Field, filter.Value)
			if err != nil {
				return ctx, err
			}
		}
	}
	return ctx, nil
}

func (f *Filters) Purge(validFields []string) {
	validFilters := make([]Filter, 0)
	for _, filter := range *f {
		if slices.Contains(validFields, filter.Field) {
			validFilters = append(validFilters, filter)
		}
	}
	*f = validFilters
}
