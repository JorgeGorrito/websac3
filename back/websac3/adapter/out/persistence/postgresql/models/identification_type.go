package models

type IdentificationType struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(32);not null"`
}
