package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type GetDegreeProgramPort interface {
	GetByID(id uint, db db.Context) (entity.DegreeProgram, error)
	GetByIDWithLang(id uint, lang string, db db.Context) (entity.DegreeProgram, error)
	GetByFilters(page, perPage uint, filters filter.Filters, db db.Context) ([]entity.DegreeProgram, int64, error)
	GetByUserIDAndFilters(page, perPage uint, userID uint, filters filter.Filters, db db.Context) ([]entity.DegreeProgram, int64, error)
}
