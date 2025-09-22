package query

import "websac3/common/paginator"

type ListCourseNaturesQuery struct {
	paginator.PaginationParams
	Filters     map[string]interface{}
	Permissions []string
}
