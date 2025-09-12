package model

type KnowledgeAreaReport struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ReportID uint   `gorm:"not null;index" json:"report_id"`
	Report   Report `gorm:"foreignKey:ReportID"`

	Name string `gorm:"type:varchar(255);not null" json:"name"`

	TotalLearnHoursExpected float32 `gorm:"not null" json:"total_learn_hours_expected"`
	TotalLearnHoursActual   float32 `gorm:"not null" json:"total_learn_hours_actual"`

	ScoreExpected float32 `gorm:"not null" json:"score_expected"`
	ScoreGot      float32 `gorm:"not null" json:"score_got"`

	TopicReports []TopicReport `gorm:"foreignKey:KnowledgeAreaReportID" json:"topic_reports"`
}

func (KnowledgeAreaReport) TableName() string {
	return "knowledge_area_reports"
}
