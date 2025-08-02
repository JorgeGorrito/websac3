package filter

import "websac3/app/port/out/persistence"

type Operator string
type Filter struct {
	Field    string
	Operator Operator
	Value    any
}

type FilterFunc func(ctx persistence.Context, field string, value any) (persistence.Context, error)
type FilterFuncMap map[Operator]FilterFunc
type Filters []Filter

func (f Filters) Apply(baseCtx persistence.Context, registryFilters FilterFuncMap) (persistence.Context, error) {
	var err error
	ctx := baseCtx
	for _, filter := range f {
		if fn, ok := registryFilters[filter.Operator]; ok {
			ctx, err = fn(ctx, filter.Field, filter.Value)
			if err != nil {
				return ctx, err
			}
		}
	}
	return ctx, nil
}
