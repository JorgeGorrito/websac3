package query

type ListReportsByDegreeProgramQuery struct {
	DegreeProgramID uint `json:"degree_program_id" validate:"required"`
}
