package request

type CreateAccessRequestRequest struct {
	Person             CreatePersonRequest `json:"person" `
	ValidationEmailURL string              `json:"validation_email_url" example:"http://localhost:3000/validate-email/token="`
	RegisterUserURL    string              `json:"register_user_url" example:"http://localhost:3000/register-user/token="`
}
