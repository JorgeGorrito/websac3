package model

import (
	"time"

	"gorm.io/gorm"
)

type DegreeProgram struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	Snies               uint           `gorm:"not null" json:"snies"`
	Name                string         `gorm:"not null" json:"name"`
	TotalCredits        uint           `json:"total_credits"`
	DurationValue       uint           `json:"duration_value"`
	DurationUnitID      uint           `gorm:"not null" json:"duration_unit_id"`
	DurationUnit        DurationUnit   `gorm:"foreignKey:DurationUnitID"`
	FormationLevelID    uint           `gorm:"not null" json:"formation_level_id"`
	FormationLevel      FormationLevel `gorm:"foreignKey:FormationLevelID"`
	ProgramFocus        string         `gorm:"varchar(255)" json:"program_focus"`
	EntryProfile        string         `gorm:"varchar(255)" json:"entry_profile"`
	GraduateProfile     string         `gorm:"varchar(255)" json:"graduate_profile"`
	ProfessionalProfile string         `gorm:"varchar(255)" json:"professional_profile"`

	Courses []Course `gorm:"foreignKey:DegreeProgramID" json:"courses"`

	ProfessionalRoles []ProfessionalRole `gorm:"many2many:degree_program_professional_roles;" json:"professional_roles"`

	CreatedBy   uint `gorm:"not null" json:"created_by"`
	UserCreator User `gorm:"foreignKey:CreatedBy"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
