package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListDurationUnitQuery struct {
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
}
