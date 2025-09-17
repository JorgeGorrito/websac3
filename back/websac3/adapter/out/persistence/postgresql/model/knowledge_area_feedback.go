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

	Comments        string  `gorm:"type:text" json:"comments"`
	SecurityGaps    string  `gorm:"type:text" json:"security_gaps"`
	Improvements    string  `gorm:"type:text" json:"improvements"`
	ComplianceLevel string  `gorm:"type:varchar(50)" json:"compliance_level"`
	AuditorRating   float32 `gorm:"type:decimal(2,1);check:auditor_rating >= 1 AND auditor_rating <= 5" json:"auditor_rating"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (KnowledgeAreaFeedback) TableName() string {
	return "knowledge_area_feedbacks"
}
