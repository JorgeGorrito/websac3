package model

import (
	"time"

	"gorm.io/gorm"
)

type ProfessionalRoleKnowledgeAreaTopic struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ProfessionalRoleKnowledgeAreaID uint                          `gorm:"not null;index" json:"professional_role_knowledge_area_id"`
	ProfessionalRoleKnowledgeArea   ProfessionalRoleKnowledgeArea `gorm:"foreignKey:ProfessionalRoleKnowledgeAreaID"`

	TopicID uint  `gorm:"not null;index" json:"topic_id"`
	Topic   Topic `gorm:"foreignKey:TopicID"`

	StudyHours uint `gorm:"not null" json:"study_hours"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ProfessionalRoleKnowledgeAreaTopic) TableName() string {
	return "professional_role_knowledge_area_topics"
}
