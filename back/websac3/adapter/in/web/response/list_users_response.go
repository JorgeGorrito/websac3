package response

import "websac3/common/paginator"

type ListUsersResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type ListUsersApiResponse struct {
	HttpStatusCode int                                `json:"http_status_code"`
	Result         *paginator.Page[ListUsersResponse] `json:"result"`
	Errors         []string                           `json:"errors"`
}
