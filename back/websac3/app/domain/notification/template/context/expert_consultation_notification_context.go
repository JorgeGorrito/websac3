package context

type ExpertConsultationNotificationContext struct {
	// Información del solicitante
	RequesterName  string `json:"requester_name"`
	RequesterEmail string `json:"requester_email"`

	// Información de la institución educativa del solicitante
	RequesterInstitutionName      string `json:"requester_institution_name"`
	RequesterInstitutionOwnership string `json:"requester_institution_ownership"`
	RequesterJobPosition          string `json:"requester_job_position"`

	// Información del programa de grado
	DegreeProgramName  string `json:"degree_program_name"`
	DegreeProgramSnies uint   `json:"degree_program_snies"`

	// Información del reporte
	ReportScore float32 `json:"report_score"`

	// Mensaje de solicitud original
	RequestMessage string `json:"request_message"`

	// Respuesta del experto
	ExpertResponse string `json:"expert_response"`

	// Estado de la consulta
	Status string `json:"status"` // "accepted" o "rejected"

	// Fechas
	CreatedAt  string `json:"created_at"`
	AnsweredAt string `json:"answered_at"`
}
