package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseTopic struct {
	CourseID uint   `gorm:"primaryKey" json:"course_id"`
	Course   Course `gorm:"foreignKey:CourseID"`
	TopicID  uint   `gorm:"primaryKey" json:"topic_id"`
	Topic    Topic  `gorm:"foreignKey:TopicID"`

	StudyHours uint `gorm:"not null" json:"study_hours"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
