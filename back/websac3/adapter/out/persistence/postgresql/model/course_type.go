package model

import "gorm.io/gorm"

type CourseType struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []CourseTypeName

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
