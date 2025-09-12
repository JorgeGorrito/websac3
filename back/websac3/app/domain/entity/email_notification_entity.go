package entity

import "time"

type EmailNotification struct {
	ID          uint
	To          string
	Subject     string
	Content     string
	CreatedAt   time.Time
	SentAt      *time.Time
	Attachments []EmailAttachment // Para envío en memoria (no se persiste)
}
