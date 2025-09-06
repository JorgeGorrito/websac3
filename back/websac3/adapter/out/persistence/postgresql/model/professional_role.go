package model

import (
	"time"

	"gorm.io/gorm"
)

type ProfessionalRole struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(128);not null" json:"name"`

	KnowledgeAreas []ProfessionalRoleKnowledgeArea `gorm:"foreignKey:ProfessionalRoleID" json:"knowledge_areas"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ProfessionalRole) TableName() string {
	return "professional_roles"
}
