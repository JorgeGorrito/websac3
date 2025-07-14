package model

import (
	"time"
)

type Role struct {
	ID          uint         `gorm:"primaryKey" `
	Name        string       `gorm:"not null" `
	Users       []User       `gorm:"foreignKey:RoleID" `
	Permissions []Permission `gorm:"many2many:role_permissions" `
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
