package response

type AnsweredExpertConsultationResponse struct {
	ID             uint   `json:"id"`
	RequesterID    uint   `json:"requester_id"`
	RequesterName  string `json:"requester_name"`
	RequesterEmail string `json:"requester_email"`

	// Información de la institución educativa del solicitante
	RequesterInstitutionSnies     uint   `json:"requester_institution_snies"`
	RequesterInstitutionName      string `json:"requester_institution_name"`
	RequesterInstitutionOwnership string `json:"requester_institution_ownership"`
	RequesterJobPosition          string `json:"requester_job_position"`

	DegreeProgramID    uint   `json:"degree_program_id"`
	DegreeProgramName  string `json:"degree_program_name"`
	DegreeProgramSnies uint   `json:"degree_program_snies"`

	ReportID    uint    `json:"report_id"`
	ReportScore float32 `json:"report_score"`

	RequestMessage *string `json:"request_message"`
	ExpertResponse *string `json:"expert_response"`

	StatusID   uint   `json:"status_id"`
	StatusName string `json:"status_name"`

	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
	AnsweredAt *string `json:"answered_at"`
}





