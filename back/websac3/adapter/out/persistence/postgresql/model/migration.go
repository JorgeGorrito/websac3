package model

import "time"

type Migration struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	ExecutedAt  time.Time `gorm:"not null"`
	Description string    `gorm:"type:text"`
}
