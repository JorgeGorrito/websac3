package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/filter"
)

type ListDegreeProgramUseCase interface {
	Execute(page, perPage uint, filters filter.Filters, userRole string, userID uint, lang string) ([]entity.DegreeProgram, int64, error)
}
