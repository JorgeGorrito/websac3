package models

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Roles       []Role `gorm:"many2many:role_permissions"`
}
