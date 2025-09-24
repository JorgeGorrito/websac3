package response

import "time"

type CreateExpertConsultationResponse struct {
	ID              uint      `json:"id" example:"1"`
	RequesterID     uint      `json:"requester_id" example:"1"`
	DegreeProgramID uint      `json:"degree_program_id" example:"1"`
	ReportID        uint      `json:"report_id" example:"1"`
	RequestMessage  *string   `json:"request_message,omitempty" example:"Necesito asesoría sobre ciberseguridad en mi programa de grado"`
	Status          string    `json:"status" example:"pending"`
	CreatedAt       time.Time `json:"created_at" example:"2024-01-15T10:30:00Z"`
}
