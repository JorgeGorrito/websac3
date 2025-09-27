package command

type CreateReportFeedbackCommand struct {
	ReportID               uint                                 `json:"report_id" validate:"required"`
	GeneralComments        string                               `json:"general_comments" validate:"required"`
	Recommendations        string                               `json:"recommendations"`
	KnowledgeAreaFeedbacks []CreateKnowledgeAreaFeedbackCommand `json:"knowledge_area_feedbacks"`
	UserID                 uint                                 `json:"-"`
}

type CreateKnowledgeAreaFeedbackCommand struct {
	KnowledgeAreaReportID uint   `json:"knowledge_area_report_id" validate:"required"`
	Comments              string `json:"comments"`
}
