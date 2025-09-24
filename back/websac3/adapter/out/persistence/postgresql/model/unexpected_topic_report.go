package model

type UnexpectedTopicReport struct {
	ID                              uint                          `gorm:"primaryKey" json:"id"`
	ReportID                        uint                          `gorm:"not null;index" json:"report_id"`
	Report                          Report                        `gorm:"foreignKey:ReportID"`
	UnexpectedKnowledgeAreaReportID uint                          `gorm:"not null;index" json:"unexpected_knowledge_area_report_id"`
	UnexpectedKnowledgeAreaReport   UnexpectedKnowledgeAreaReport `gorm:"foreignKey:UnexpectedKnowledgeAreaReportID"`
	TopicID                         uint                          `gorm:"not null" json:"topic_id"`
	Topic                           Topic                         `gorm:"foreignKey:TopicID"`
	LearnHoursActual                float32                       `gorm:"not null" json:"learn_hours_actual"`
	KnowledgeAreaID                 uint                          `gorm:"not null" json:"knowledge_area_id"`
}

func (UnexpectedTopicReport) TableName() string {
	return "unexpected_topic_reports"
}
