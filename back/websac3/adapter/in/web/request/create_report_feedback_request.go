package request

type CreateReportFeedbackRequest struct {
	ReportID               uint                                 `json:"report_id" binding:"required"`
	GeneralComments        string                               `json:"general_comments" binding:"required"`
	Recommendations        string                               `json:"recommendations"`
	AuditorRating          float32                              `json:"auditor_rating" binding:"required,min=1,max=5"`
	KnowledgeAreaFeedbacks []CreateKnowledgeAreaFeedbackRequest `json:"knowledge_area_feedbacks"`

	Permissions []string `json:"permissions"`
}

type CreateKnowledgeAreaFeedbackRequest struct {
	KnowledgeAreaReportID uint    `json:"knowledge_area_report_id" binding:"required"`
	Comments              string  `json:"comments"`
	SecurityGaps          string  `json:"security_gaps"`
	Improvements          string  `json:"improvements"`
	ComplianceLevel       string  `json:"compliance_level"`
	AuditorRating         float32 `json:"auditor_rating" binding:"required,min=1,max=5"`
}
