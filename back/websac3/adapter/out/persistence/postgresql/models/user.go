package models

import "time"

type User struct {
	ID            uint       `gorm:"primaryKey" mapper:"userID"`
	Email         string     `gorm:"type:varchar(128); uniqueIndex" mapper:"userEmail"`
	PasswordHash  string     `gorm:"type:varchar(256); null" mapper:"userPasswordHash"`
	RoleID        *uint      `gorm:"default:null" mapper:"roleID"`
	Role          *Role      `gorm:"foreignKey:RoleID" mapper:"role"`
	DeactivatedAt *time.Time `gorm:"default:null" mapper:"userDeletedAt"`
	CreatedAt     time.Time  `mapper:"userCreatedAt"`
}
