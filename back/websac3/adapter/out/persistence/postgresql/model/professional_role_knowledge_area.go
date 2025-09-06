package model

import (
	"time"

	"gorm.io/gorm"
)

type ProfessionalRoleKnowledgeArea struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ProfessionalRoleID uint             `gorm:"not null;index" json:"professional_role_id"`
	ProfessionalRole   ProfessionalRole `gorm:"foreignKey:ProfessionalRoleID"`

	KnowledgeAreaID uint          `gorm:"not null;index" json:"knowledge_area_id"`
	KnowledgeArea   KnowledgeArea `gorm:"foreignKey:KnowledgeAreaID"`

	PriorityWeight float32 `gorm:"not null" json:"priority_weight"`

	Topics []ProfessionalRoleKnowledgeAreaTopic `gorm:"foreignKey:ProfessionalRoleKnowledgeAreaID" json:"topics"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ProfessionalRoleKnowledgeArea) TableName() string {
	return "professional_role_knowledge_areas"
}
