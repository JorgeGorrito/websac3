package filter

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/filter"
)

const (
	EqualOperator    filter.Operator = "eq"
	ContainsOperator filter.Operator = "cont"
)

var filtersRegistry filter.FilterFuncMap = map[filter.Operator]filter.FilterFunc{
	EqualOperator: func(ctx persistence.Context, field string, value any) (persistence.Context, error) {
		dbCtx, ok := ctx.(*db.Context)
		if !ok {
			return ctx, errors.New("db context cast error")
		}

		if err := dbCtx.DB().
			Where(field+" = ?", value).
			Error; err != nil {
			return ctx, err
		}

		return ctx, nil
	},
	ContainsOperator: func(ctx persistence.Context, field string, value any) (persistence.Context, error) {
		dbCtx, ok := ctx.(*db.Context)
		if !ok {
			return ctx, errors.New("db context cast error")
		}

		if err := dbCtx.DB().
			Where(field+" LIKE ?", "%"+value.(string)+"%").
			Error; err != nil {
			return ctx, err
		}

		return ctx, nil
	},
}
