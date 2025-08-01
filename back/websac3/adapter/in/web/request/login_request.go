package request

type LoginRequest struct {
	Email    string `json:"email" example:"admin@websac3.com"`
	Password string `json:"password" example:"admin123"`
}
