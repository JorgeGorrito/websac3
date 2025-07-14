package command

type CreateAccessRequestCommand struct {
	Person        CreatePersonCommand `validations:"required"`
	RedirectUrlTo string              `validations:"required"`
}
