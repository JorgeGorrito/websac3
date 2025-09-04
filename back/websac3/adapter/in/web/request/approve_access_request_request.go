package request

type ApproveAccessRequestRequest struct {
	RoleID uint `json:"role_id" binding:"required" example:"1"`
}
