package model

type InstitutionalCategory struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(128);not null"`
}
