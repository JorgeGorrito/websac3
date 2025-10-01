package request

type ChangePasswordRequest struct {
	CurrentPassword    string `json:"current_password" binding:"required,min=8" example:"currentPassword123"`
	NewPassword        string `json:"new_password" binding:"required,min=8" example:"newPassword456"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required,min=8,eqfield=NewPassword" example:"newPassword456"`
}
