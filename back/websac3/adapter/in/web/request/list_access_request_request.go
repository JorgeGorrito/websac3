package request

import (
	"websac3/common/paginator"
)

type ListAccessRequestRequest struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
}
