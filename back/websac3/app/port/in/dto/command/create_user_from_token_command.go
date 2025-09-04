package command

type CreateUserFromTokenCommand struct {
	CreateUserToken string `validations:"required"`
	Password        string `validations:"required;min=8"`
	ConfirmPassword string `validations:"required;min=8"`
}
