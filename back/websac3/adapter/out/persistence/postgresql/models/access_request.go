package models

import (
	"time"
)

type AccessRequest struct {
	ID uint `gorm:"primaryKey" mapper:"accessRequestID"`

	UpdatedBy *uint `gorm:"null"`
	User      User  `gorm:"foreignKey:UpdatedBy"`

	ApplicantID uint   `gorm:"not null" mapper:"personID"`
	Applicant   Person `gorm:"foreignKey:ApplicantID"`

	StatusID  uint      `gorm:"not null" mapper:"statusID"`
	Status    Status    `gorm:"foreignKey:StatusID"`
	CreatedAt time.Time `mapper:"accessRequestcreatedAt"`
	UpdatedAt time.Time `mapper:"accessRequestupdatedAt"`
	DeleteAt  time.Time `mapper:"accessRequestdeleteAt"`
}
