package models

type Municipality struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100);not null"`
}
