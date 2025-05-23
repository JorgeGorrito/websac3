package persistence

import "websac3/app/domain/entity"

type GetAccessRequestPort interface {
	GetLastCreatedPersonIdentificationNumber(identificationNumber string, tx Transaction) (entity.AccessRequest, error)
}
