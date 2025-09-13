package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListRejectedAccessRequestQuery struct {
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
	UserID           uint
	Permissions      []string
}
