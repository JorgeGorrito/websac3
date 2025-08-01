package command

type LoginCommand struct {
	Email    string `validations:"required;email"`
	Password string `validations:"required"`
}
