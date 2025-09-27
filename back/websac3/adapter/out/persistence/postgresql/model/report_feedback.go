package model

import (
	"time"

	"gorm.io/gorm"
)

type ReportFeedback struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	ReportID uint   `gorm:"not null;index" json:"report_id"`
	Report   Report `gorm:"foreignKey:ReportID"`

	AuditorID uint `gorm:"not null;index" json:"auditor_id"`
	Auditor   User `gorm:"foreignKey:AuditorID"`

	GeneralComments string `gorm:"type:text" json:"general_comments"`
	Recommendations string `gorm:"type:text" json:"recommendations"`

	KnowledgeAreaFeedbacks []KnowledgeAreaFeedback `gorm:"foreignKey:ReportFeedbackID" json:"knowledge_area_feedbacks"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ReportFeedback) TableName() string {
	return "report_feedbacks"
}
