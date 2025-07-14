package persistence

import "websac3/app/domain/entity"

type GetPersonPort interface {
	GetByIdentificationNumber(identificationNumber string, db Context) (entity.Person, error)
}
