package response

type ApiResponse[T any] struct {
	HttpStatusCode int
	Result         T
	Error          string
}

func (r *ApiResponse[T]) ToResponseFormat() map[string]any {
	return map[string]any{
		"result": r.Result,
		"error":  r.Error,
	}
}
