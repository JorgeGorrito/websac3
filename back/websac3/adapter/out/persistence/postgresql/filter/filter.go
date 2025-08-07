package filter

import (
	"errors"
	"strings"
	"websac3/adapter/out/persistence/postgresql/db"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

func getFieldFormatted(field string) string {
	fieldSplit := strings.Split(field, ".")
	if len(fieldSplit) > 1 {
		var fieldFormatted string = "\""
		var numberParts int = len(fieldSplit)
		var i = 0
		var iNotRequiredMapped = numberParts - 2
		for _, part := range fieldSplit[:numberParts-1] {
			fieldFormatted += part
			if numberParts > 2 && i != iNotRequiredMapped {
				fieldFormatted += "__"
			}
			i++
		}

		fieldFormatted += "\"." + fieldSplit[numberParts-1]
		return fieldFormatted
	}
	return field
}

var FiltersRegistry filter.FilterFuncMap = map[filter.Operator]filter.FilterFunc{
	filter.EqualOperator: func(ctx _db.Context, field string, value any) (_db.Context, error) {
		dbCtx, ok := ctx.(*db.Context)
		if !ok {
			return ctx, errors.New("db context cast error")
		}

		fieldFormatted := getFieldFormatted(field)
		dbCtx.DBSet(dbCtx.DB().Where(fieldFormatted+" = ?", value))

		return dbCtx, nil
	},

	filter.ContainsOperator: func(ctx _db.Context, field string, value any) (_db.Context, error) {
		dbCtx, ok := ctx.(*db.Context)
		if !ok {
			return ctx, errors.New("db context cast error")
		}

		fieldFormatted := getFieldFormatted(field)
		dbCtx.DBSet(dbCtx.DB().Where(fieldFormatted+" ILIKE ?", "%"+value.(string)+"%"))
		return dbCtx, nil
	},
}
