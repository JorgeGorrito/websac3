package models

import (
	"time"
)

type Role struct {
	ID          uint         `gorm:"primaryKey" mapper:"roleID"`
	Name        string       `gorm:"not null" mapper:"roleName"`
	Users       []User       `gorm:"foreignKey:RoleID" mapper:"users"`
	Permissions []Permission `gorm:"many2many:role_permissions" mapper:"permissions"`
	CreatedAt   time.Time    `mapper:"userCreatedAt"`
	UpdatedAt   time.Time    `mapper:"userUpdatedAt"`
	DeletedAt   time.Time    `mapper:"userDeleteAt"`
}
