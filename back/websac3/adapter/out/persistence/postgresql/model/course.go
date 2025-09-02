package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID                          uint         `gorm:"primaryKey" json:"id"`
	Name                        string       `gorm:"type:varchar(255);not null" json:"name"`
	Code                        string       `gorm:"type:varchar(20);not null" json:"code"`
	Credits                     uint         `gorm:"not null" json:"credits"`
	PeriodNumber                uint         `gorm:"not null" json:"period_number"`
	NatureID                    uint         `gorm:"not null;index" json:"nature_id"`
	Nature                      CourseNature `gorm:"foreignKey:NatureID"`
	TypeID                      uint         `gorm:"not null;index" json:"type_id"`
	Type                        CourseType   `gorm:"foreignKey:TypeID"`
	IsCybersecurity             bool         `gorm:"default:false" json:"is_cybersecurity"`
	ContainsCybersecurityTopics bool         `gorm:"default:false" json:"contains_cybersecurity_topics"`

	// Relación con DegreeProgram
	DegreeProgramID uint          `gorm:"not null;index" json:"degree_program_id"`
	DegreeProgram   DegreeProgram `gorm:"foreignKey:DegreeProgramID"`

	// Relación muchos a muchos con Topic (Temáticas) con información de tiempo de estudio
	CourseTopics []CourseTopic `json:"course_topics"`

	// Relación con el usuario que creó el curso
	CreatedBy   uint `gorm:"not null" json:"created_by"`
	UserCreator User `gorm:"foreignKey:CreatedBy"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName especifica el nombre de la tabla
func (Course) TableName() string {
	return "courses"
}

// AfterMigrate crea el índice compuesto único después de la migración
func (Course) AfterMigrate(tx *gorm.DB) error {
	// Crear índice compuesto único para code + degree_program_id
	return tx.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_courses_code_degree_program_unique 
		ON courses (code, degree_program_id) 
		WHERE deleted_at IS NULL
	`).Error
}
