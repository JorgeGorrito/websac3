package notification

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type SendMailPort interface {
	Send(notification *entity.EmailNotification, ctx db.Context) error
}
