package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetPersonPort interface {
	GetByIdentificationNumber(identificationNumber string, db db.Context) (entity.Person, error)
}
