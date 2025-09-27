package request

type CreateReportFeedbackRequest struct {
	ReportID               uint                                 `json:"report_id" binding:"required"`
	GeneralComments        string                               `json:"general_comments" binding:"required"`
	Recommendations        string                               `json:"recommendations"`
	KnowledgeAreaFeedbacks []CreateKnowledgeAreaFeedbackRequest `json:"knowledge_area_feedbacks"`
}

type CreateKnowledgeAreaFeedbackRequest struct {
	KnowledgeAreaReportID uint   `json:"knowledge_area_report_id" binding:"required"`
	Comments              string `json:"comments"`
}
