package command

type CreateUserCommand struct {
	Email string `validate:"required;email" mapper:"userEmail"`
}
