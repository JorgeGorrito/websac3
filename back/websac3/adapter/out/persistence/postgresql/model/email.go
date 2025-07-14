package model

import "time"

type Email struct {
	ID        uint       `gorm:"primaryKey" `
	ToEmail   string     `gorm:"not null" `
	Subject   string     `gorm:"not null" `
	Body      string     `gorm:"not null" `
	CreatedAt time.Time  `gorm:"not null" `
	SentAt    *time.Time `gorm:"null" `
}
