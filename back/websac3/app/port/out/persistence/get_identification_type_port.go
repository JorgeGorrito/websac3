package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type GetIdentificationTypePort interface {
	GetByFilters(page, perPage uint, filters filter.Filters, db db.Context) ([]entity.IdentificationType, int64, error)
}
