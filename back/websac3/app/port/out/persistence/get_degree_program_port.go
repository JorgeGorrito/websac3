package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type GetDegreeProgramPort interface {
	GetByFilters(page, perPage uint, filters filter.Filters, db db.Context) ([]entity.DegreeProgram, int64, error)
	GetByIDAndFilters(page, perPage uint, userID uint, filters filter.Filters, db db.Context) ([]entity.DegreeProgram, int64, error)
}
