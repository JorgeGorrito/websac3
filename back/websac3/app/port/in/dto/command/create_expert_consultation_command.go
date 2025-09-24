package command

type CreateExpertConsultationCommand struct {
	RequesterID     uint    `validations:"required"`
	DegreeProgramID uint    `validations:"required"`
	ReportID        uint    `validations:"required"`
	RequestMessage  *string `validations:"optional"`
}
