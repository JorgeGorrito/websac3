package request

import "websac3/common/paginator"

type ListTopicRequest struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
}
