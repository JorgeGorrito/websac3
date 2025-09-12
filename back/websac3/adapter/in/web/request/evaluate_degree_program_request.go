package request

type EvaluateDegreeProgramRequest struct {
	DegreeProgramID    uint `json:"degree_program_id" validations:"required"`
	ProfessionalRoleID uint `json:"professional_role_id" validations:"required"`
}
