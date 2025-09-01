package model

import "gorm.io/gorm"

type CourseNature struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []CourseNatureName

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
