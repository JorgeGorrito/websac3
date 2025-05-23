package models

type Ownership struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(16);not null"`
}
