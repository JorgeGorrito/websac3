package filter

import (
	"slices"
	"websac3/app/port/out/persistence/filter"
)

type Params map[string]map[string]string

func Transform(params Params, operators []filter.Operator) filter.Filters {
	var filters filter.Filters
	for field, filterData := range params {
		for operator, value := range filterData {
			if value == "" || !slices.Contains(operators, filter.Operator(operator)) {
				continue
			}
			filters = append(filters, filter.Filter{
				Field:    field,
				Operator: filter.Operator(operator),
				Value:    value,
			})
		}
	}
	return filters
}
