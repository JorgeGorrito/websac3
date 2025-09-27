package model

import (
	"time"

	"gorm.io/gorm"
)

type KnowledgeAreaFeedback struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	ReportFeedbackID uint           `gorm:"not null;index" json:"report_feedback_id"`
	ReportFeedback   ReportFeedback `gorm:"foreignKey:ReportFeedbackID"`

	KnowledgeAreaReportID uint                `gorm:"not null;index" json:"knowledge_area_report_id"`
	KnowledgeAreaReport   KnowledgeAreaReport `gorm:"foreignKey:KnowledgeAreaReportID"`

	Comments string `gorm:"type:text" json:"comments"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (KnowledgeAreaFeedback) TableName() string {
	return "knowledge_area_feedbacks"
}
