package command

type DeactivateUserCommand struct {
	UserID      uint     `json:"user_id" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}
