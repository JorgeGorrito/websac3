package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListTopicQuery struct {
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
	Permissions      []string
}
