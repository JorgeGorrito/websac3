package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListDegreeProgramQuery struct {
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
	Permissions      []string
	UserRole         string
	UserID           uint
}
