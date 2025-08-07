package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type UpdatePersonPort interface {
	UpdateById(person *entity.Person, personID uint, db db.Context) error
}
