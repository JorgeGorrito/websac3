package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListReportsPendingFeedbackQuery struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
	UserID           uint                       `json:"-"`
	Filters          filter.Params              `json:"filters"`
	SortBy           string                     `json:"sort_by"`           // Campo por el cual ordenar
	SortOrder        string                     `json:"sort_order"`        // "asc" o "desc"
	Permissions      []string                   `json:"permissions"`
}
