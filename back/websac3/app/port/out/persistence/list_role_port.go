package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListRolePort interface {
	GetByFilters(page, perPage uint, filters filter.Filters, db db.Context) ([]entity.Role, int64, error)
}
