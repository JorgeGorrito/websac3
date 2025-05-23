package persistence

import "websac3/app/domain/entity"

type GetPersonPort interface {
	GetByIdentificationNumber(identificationNumber string, tx Transaction) (entity.Person, error)
}
