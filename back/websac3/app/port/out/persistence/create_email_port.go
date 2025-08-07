package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateEmailPort interface {
	Create(emailToCreate *entity.EmailNotification, db db.Context) error
}
