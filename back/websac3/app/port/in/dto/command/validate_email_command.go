package command

type ValidateEmailCommand struct {
	ValidationToken string `validations:"required"`
}
