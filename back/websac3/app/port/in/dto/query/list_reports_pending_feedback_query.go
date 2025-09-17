package query

import "websac3/common/paginator"

type ListReportsPendingFeedbackQuery struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
	UserID           uint                       `json:"-"`

	Permissions []string `json:"permissions"`
}
