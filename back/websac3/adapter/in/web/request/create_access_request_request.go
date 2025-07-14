package request

type CreateAccessRequestRequest struct {
	Person        CreatePersonRequest `json:"person" `
	RedirectUrlTo string              `json:"redirect_url_to" `
}
