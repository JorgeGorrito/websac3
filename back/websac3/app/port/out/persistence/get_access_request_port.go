package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type GetAccessRequestPort interface {
	GetAuthenticatedEmailByFilters(page, perPage uint, filters filter.Filters, db db.Context) ([]entity.AccessRequest, int64, error)
	GetLastCreatedByIdentificationAndEmail(identificationTypeID uint, identificationNumber string, email string, db db.Context) (entity.AccessRequest, error)
	GetUnvalidatedEmailByToken(validationToken string, db db.Context) (entity.AccessRequest, error)
}
