package response

type ReportFeedbackResponse struct {
	ID                     uint                            `json:"id"`
	ReportID               uint                            `json:"report_id"`
	Report                 ReportSummaryResponse           `json:"report"`
	Auditor                UserSummaryResponse             `json:"auditor"`
	GeneralComments        string                          `json:"general_comments"`
	Recommendations        string                          `json:"recommendations"`
	AuditorRating          float32                         `json:"auditor_rating"`
	KnowledgeAreaFeedbacks []KnowledgeAreaFeedbackResponse `json:"knowledge_area_feedbacks"`
	CreatedAt              string                          `json:"created_at"`
	UpdatedAt              string                          `json:"updated_at"`
}

type KnowledgeAreaFeedbackResponse struct {
	ID                    uint    `json:"id"`
	KnowledgeAreaReportID uint    `json:"knowledge_area_report_id"`
	KnowledgeAreaName     string  `json:"knowledge_area_name"`
	Comments              string  `json:"comments"`
	SecurityGaps          string  `json:"security_gaps"`
	Improvements          string  `json:"improvements"`
	ComplianceLevel       string  `json:"compliance_level"`
	AuditorRating         float32 `json:"auditor_rating"`
}

type ReportSummaryResponse struct {
	ID            uint                         `json:"id"`
	Score         float32                      `json:"score"`
	DegreeProgram DegreeProgramSummaryResponse `json:"degree_program"`
	CreatedAt     string                       `json:"created_at"`
}

type DegreeProgramSummaryResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Snies uint   `json:"snies"`
}

type UserSummaryResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ReportResponse struct {
	ID                         uint                         `json:"id"`
	Score                      float32                      `json:"score"`
	DegreeProgram              DegreeProgramSummaryResponse `json:"degree_program"`
	ProfessionalRole           ProfessionalRoleResponse     `json:"professional_role"`
	HigherEducationInstitution InstitutionSummaryResponse   `json:"higher_education_institution"`
	CreatedAt                  string                       `json:"created_at"`
	HasFeedback                bool                         `json:"has_feedback"`
}

type InstitutionSummaryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PaginatedResponse[T any] struct {
	Data       []T  `json:"data"`
	Total      uint `json:"total"`
	Page       uint `json:"page"`
	Limit      uint `json:"limit"`
	TotalPages uint `json:"total_pages"`
}
