package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ProfessionalRoleID uint             `gorm:"not null;index" json:"professional_role_id"`
	ProfessionalRole   ProfessionalRole `gorm:"foreignKey:ProfessionalRoleID"`

	DegreeProgramID uint          `gorm:"not null;index" json:"degree_program_id"`
	DegreeProgram   DegreeProgram `gorm:"foreignKey:DegreeProgramID"`

	Score float32 `gorm:"not null" json:"score"`

	KnowledgeAreaReports []KnowledgeAreaReport `gorm:"foreignKey:ReportID" json:"knowledge_area_reports"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Report) TableName() string {
	return "reports"
}
