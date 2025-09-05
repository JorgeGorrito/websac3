package model

import (
	"time"
)

type AccessRequest struct {
	ID uint `gorm:"primaryKey"`

	ApplicantID uint   `gorm:"not null" `
	Applicant   Person `gorm:"foreignKey:ApplicantID"`

	ValidationEmailCode      string `gorm:"not null; type:uuid" `
	ValidationCreateUserCode string `gorm:"not null; type:uuid" `

	ValidationEmailURL string `gorm:"not null; type:varchar(255)" `
	CreateUserURL      string `gorm:"not null; type:varchar(255)" `

	ApprovedEmailID *uint  `gorm:"null" `
	ApprovedEmail   *Email `gorm:"foreignKey:ApprovedEmailID" `

	VerificationEmailID *uint  `gorm:"null" `
	VerificationEmail   *Email `gorm:"foreignKey:VerificationEmailID" `

	StatusID uint   `gorm:"not null" `
	Status   Status `gorm:"foreignKey:StatusID" `

	ApprovedRoleID *uint `gorm:"null"`
	ApprovedRole   Role  `gorm:"foreignKey:ApprovedRoleID"`

	Lang string `gorm:"not null; type:varchar(5)"`

	IsVerified bool `gorm:"not null; default:false" `
	CreatedAt  time.Time
}
