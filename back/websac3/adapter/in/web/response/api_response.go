package response

type ApiResponse[T any] struct {
	HttpStatusCode int
	Result         T
}

func (r *ApiResponse[T]) ToResponseFormat() map[string]T {
	return map[string]T{
		"result": r.Result,
	}
}
