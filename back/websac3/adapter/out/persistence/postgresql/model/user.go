package model

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey" `
	Email        string `gorm:"type:varchar(128); uniqueIndex" `
	PasswordHash string `gorm:"type:varchar(256); null" json:"password_hash"`

	RoleID uint `gorm:"not null" json:"role_id"`
	Role   Role `gorm:"foreignKey:RoleID" `

	PersonID uint   `gorm:"not null" json:"person_id"`
	Person   Person `gorm:"foreignKey:PersonID" `

	DeactivatedAt *time.Time `gorm:"default:null" json:"deactivated_at"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
}
