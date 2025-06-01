package response

import (
	"errors"
)

var (
	failedToMarshalResult = errors.New("failed to marshal result")
)

type ApiResponse[T any] struct {
	HttpStatusCode int `json:"httpStatusCode"`
	Result         T   `json:"result"`
}

func (r *ApiResponse[T]) ToResponseFormat() map[string]T {
	return map[string]T{
		"result": r.Result,
	}
}
