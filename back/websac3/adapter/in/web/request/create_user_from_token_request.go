package request

type CreateUserFromTokenRequest struct {
	CreateUserToken string `json:"create_user_token" binding:"required" example:"uuid-token-here"`
	Password        string `json:"password" binding:"required,min=8" example:"password123"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password" example:"password123"`
}
