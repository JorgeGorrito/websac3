package request

import "websac3/common/paginator"

type ListHigherEducationInstitutionRequest struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
}
