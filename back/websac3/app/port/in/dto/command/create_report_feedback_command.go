package command

type CreateReportFeedbackCommand struct {
	ReportID               uint                                 `json:"report_id" validate:"required"`
	GeneralComments        string                               `json:"general_comments" validate:"required"`
	Recommendations        string                               `json:"recommendations"`
	AuditorRating          float32                              `json:"auditor_rating" validate:"required,min=1,max=5"`
	KnowledgeAreaFeedbacks []CreateKnowledgeAreaFeedbackCommand `json:"knowledge_area_feedbacks"`
	UserID                 uint                                 `json:"-"`
	Permissions            []string                             `json:"permissions"`
}

type CreateKnowledgeAreaFeedbackCommand struct {
	KnowledgeAreaReportID uint    `json:"knowledge_area_report_id" validate:"required"`
	Comments              string  `json:"comments"`
	SecurityGaps          string  `json:"security_gaps"`
	Improvements          string  `json:"improvements"`
	ComplianceLevel       string  `json:"compliance_level"`
	AuditorRating         float32 `json:"auditor_rating" validate:"required,min=1,max=5"`
}
