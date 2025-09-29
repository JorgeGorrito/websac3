package response

type UserExpertConsultationResponse struct {
	ID             uint   `json:"id"`
	RequesterID    uint   `json:"requester_id"`
	RequesterName  string `json:"requester_name"`
	RequesterEmail string `json:"requester_email"`

	DegreeProgramID    uint   `json:"degree_program_id"`
	DegreeProgramName  string `json:"degree_program_name"`
	DegreeProgramSnies uint   `json:"degree_program_snies"`

	ReportID    uint    `json:"report_id"`
	ReportScore float32 `json:"report_score"`

	RequestMessage *string `json:"request_message"`
	ExpertResponse *string `json:"expert_response"`

	ExpertID    *uint   `json:"expert_id"`
	ExpertName  *string `json:"expert_name"`
	ExpertEmail *string `json:"expert_email"`

	StatusID   uint   `json:"status_id"`
	StatusName string `json:"status_name"`

	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
	AnsweredAt *string `json:"answered_at"`
	ClosedAt   *string `json:"closed_at"`
}
