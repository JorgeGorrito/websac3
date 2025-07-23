package model

type Module struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(128);not null"`
	Description string `gorm:"type:varchar(256); not null"`
}
