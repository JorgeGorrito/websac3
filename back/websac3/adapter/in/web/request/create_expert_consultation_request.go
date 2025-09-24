package request

type CreateExpertConsultationRequest struct {
	RequesterID     uint    `json:"requester_id" example:"1" binding:"required"`
	DegreeProgramID uint    `json:"degree_program_id" example:"1" binding:"required"`
	ReportID        uint    `json:"report_id" example:"1" binding:"required"`
	RequestMessage  *string `json:"request_message,omitempty" example:"Necesito asesoría sobre ciberseguridad en mi programa de grado"`
}
