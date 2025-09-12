package model

type TopicReport struct {
	ID uint `gorm:"primaryKey" json:"id"`

	KnowledgeAreaReportID uint                `gorm:"not null;index" json:"knowledge_area_report_id"`
	KnowledgeAreaReport   KnowledgeAreaReport `gorm:"foreignKey:KnowledgeAreaReportID"`

	TopicID uint  `gorm:"not null;index" json:"topic_id"`
	Topic   Topic `gorm:"foreignKey:TopicID"`

	Name               string  `gorm:"type:varchar(255);not null" json:"name"`
	LearnHoursExpected float32 `gorm:"not null" json:"learn_hours_expected"`
	LearnHoursActual   float32 `gorm:"not null" json:"learn_hours_actual"`
}

func (TopicReport) TableName() string {
	return "topic_reports"
}
