package command

import "websac3/common/validator"

type CreateUserCommand struct {
	Email string `validate:"required;email" mapper:"userEmail"`
}

func (c *CreateUserCommand) Validate() error {
	return validator.ValidateFields(c)
}
