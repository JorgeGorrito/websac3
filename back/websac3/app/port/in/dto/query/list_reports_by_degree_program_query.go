package query

import "websac3/common/paginator"

type ListReportsByDegreeProgramQuery struct {
	DegreeProgramID  uint                       `json:"degree_program_id" validate:"required"`
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
}
