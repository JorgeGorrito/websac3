package model

import (
	"time"

	"gorm.io/gorm"
)

type DegreeProgramProfessionalRole struct {
	ID                 uint `gorm:"primaryKey" json:"id"`
	DegreeProgramID    uint `gorm:"not null;index" json:"degree_program_id"`
	ProfessionalRoleID uint `gorm:"not null;index" json:"professional_role_id"`

	DegreeProgram    DegreeProgram    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ProfessionalRole ProfessionalRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (DegreeProgramProfessionalRole) TableName() string {
	return "degree_program_professional_roles"
}
