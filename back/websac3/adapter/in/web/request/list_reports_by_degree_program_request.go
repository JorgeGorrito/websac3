package request

type ListReportsByDegreeProgramRequest struct {
	DegreeProgramID uint `json:"degree_program_id" validate:"required"`
}
