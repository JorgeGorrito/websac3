package notification

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
)

type SendMailPort interface {
	Send(notification *entity.EmailNotification, ctx persistence.Context) error
}
