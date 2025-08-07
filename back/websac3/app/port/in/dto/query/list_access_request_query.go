package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListAccessRequestQuery struct {
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
}
