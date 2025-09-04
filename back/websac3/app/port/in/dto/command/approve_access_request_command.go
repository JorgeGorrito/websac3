package command

type ApproveAccessRequestCommand struct {
	UserID          uint     `validations:"required"`
	AccessRequestID uint     `validations:"required"`
	RoleID          uint     `validations:"required"`
	Permissions     []string `validations:"required"`
}
