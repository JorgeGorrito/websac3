package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListUsersQuery struct {
	UserID           uint
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
	Permissions      []string
}
