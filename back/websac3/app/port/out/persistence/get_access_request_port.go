package persistence

import "websac3/app/domain/entity"

type GetAccessRequestPort interface {
	GetLastCreatedByIdentificationAndEmail(identificationTypeID uint, identificationNumber string, email string, db Context) (entity.AccessRequest, error)
	GetUnvalidatedEmailByToken(validationToken string, db Context) (entity.AccessRequest, error)
}
