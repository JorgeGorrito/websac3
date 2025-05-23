package request

type CreateUserRequest struct {
	Email string `json:"email" binding:"required" mapper:"userEmail"`
}
