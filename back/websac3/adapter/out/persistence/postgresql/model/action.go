package model

type Action struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(128);unique;not null"`
}
