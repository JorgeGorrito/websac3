package query

import "websac3/common/paginator"

type ListCourseTypesQuery struct {
	paginator.PaginationParams
	Filters     map[string]interface{}
	Permissions []string
}
