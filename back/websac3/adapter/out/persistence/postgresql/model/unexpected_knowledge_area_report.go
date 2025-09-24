package model

type UnexpectedKnowledgeAreaReport struct {
	ID              uint                    `gorm:"primaryKey" json:"id"`
	ReportID        uint                    `gorm:"not null;index" json:"report_id"`
	Report          Report                  `gorm:"foreignKey:ReportID"`
	Name            string                  `gorm:"type:varchar(255);not null" json:"name"`
	Lang            string                  `gorm:"type:varchar(2);not null" json:"lang"`
	TotalLearnHours float32                 `gorm:"not null" json:"total_learn_hours"`
	TopicReports    []UnexpectedTopicReport `gorm:"foreignKey:UnexpectedKnowledgeAreaReportID" json:"topic_reports"`
}

func (UnexpectedKnowledgeAreaReport) TableName() string {
	return "unexpected_knowledge_area_reports"
}
