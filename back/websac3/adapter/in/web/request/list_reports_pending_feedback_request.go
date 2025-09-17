package request

import "websac3/common/paginator"

type ListReportsPendingFeedbackRequest struct {
	paginator.PaginationParams

	Permissions []string `json:"permissions"`
}
