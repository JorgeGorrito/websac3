package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListApprovedAccessRequestPort interface {
	GetApprovedByFilters(page, perPage uint, filters filter.Filters, db db.Context) ([]entity.AccessRequest, int64, error)
}
