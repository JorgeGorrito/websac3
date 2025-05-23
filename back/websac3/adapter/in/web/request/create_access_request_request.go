package request

type CreateAccessRequestRequest struct {
	Person CreatePersonRequest `json:"person" binding:"required" mapper:"person"`
}
