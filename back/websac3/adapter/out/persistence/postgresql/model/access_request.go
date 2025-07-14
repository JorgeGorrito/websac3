package model

import (
	"time"
)

type AccessRequest struct {
	ID uint `gorm:"primaryKey" `

	ApplicantID uint   `gorm:"not null" `
	Applicant   Person `gorm:"foreignKey:ApplicantID"`

	ValidationCode string `gorm:"not null; type:uuid" `

	VerificationEmailID *uint  `gorm:"null" `
	VerificationEmail   *Email `gorm:"foreignKey:VerificationEmailID" `

	StatusID uint   `gorm:"not null" `
	Status   Status `gorm:"foreignKey:StatusID" `

	IsVerified bool `gorm:"not null; default:false" `
	CreatedAt  time.Time
}
