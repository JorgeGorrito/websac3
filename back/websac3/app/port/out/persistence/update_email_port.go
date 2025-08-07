package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type UpdateEmailPort interface {
	Update(email *entity.EmailNotification, db db.Context) error
}
