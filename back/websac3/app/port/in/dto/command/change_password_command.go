package command

type ChangePasswordCommand struct {
	UserID             uint
	CurrentPassword    string `validations:"required;min=8"`
	NewPassword        string `validations:"required;min=8"`
	ConfirmNewPassword string `validations:"required;min=8"`
	Permissions        []string
}
