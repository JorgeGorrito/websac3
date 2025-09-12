package command

type EvaluateDegreeProgramCommand struct {
	DegreeProgramID    uint `validations:"required"`
	ProfessionalRoleID uint `validations:"required"`
	UserID             uint `validations:"required"`
}
