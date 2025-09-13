package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListProfessionalRoleQuery struct {
	PaginationParams paginator.PaginationParams
	Filters          filter.Params
}
