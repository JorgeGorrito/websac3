package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreatePersonPort interface {
	Create(person *entity.Person, db db.Context) error
}
