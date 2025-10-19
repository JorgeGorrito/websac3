package query

import (
	commonfilter "websac3/common/filter"
	"websac3/common/paginator"
)

type ListAnsweredExpertConsultationsQuery struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
	UserID           uint                       `json:"-"`
	Filters          commonfilter.Params        `json:"filters"`
	Permissions      []string                   `json:"permissions"`
}





