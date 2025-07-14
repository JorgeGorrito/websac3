package model

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey" `
	Email        string `gorm:"type:varchar(128); uniqueIndex" `
	PasswordHash string `gorm:"type:varchar(256); null" `

	RoleID *uint `gorm:"default:null" `
	Role   *Role `gorm:"foreignKey:RoleID" `

	PersonID uint   `gorm:"not null" `
	Person   Person `gorm:"foreignKey:PersonID" `

	DeactivatedAt *time.Time `gorm:"default:null" `
	CreatedAt     time.Time
}
