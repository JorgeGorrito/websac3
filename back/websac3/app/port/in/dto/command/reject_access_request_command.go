package command

type RejectAccessRequestCommand struct {
	UserID          uint     `validations:"required"`
	AccessRequestID uint     `validations:"required"`
	Permissions     []string `validations:"required"`
}
