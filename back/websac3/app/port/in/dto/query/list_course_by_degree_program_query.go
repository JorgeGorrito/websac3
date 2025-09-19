package query

import (
	"websac3/common/paginator"
)

type ListCourseByDegreeProgramQuery struct {
	paginator.PaginationParams
	DegreeProgramID uint
	Filters         map[string]interface{}
	Permissions     []string
}
