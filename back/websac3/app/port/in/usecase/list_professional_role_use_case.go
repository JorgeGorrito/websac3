package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/filter"
)

type ListProfessionalRoleUseCase interface {
	Execute(page, perPage uint, filters filter.Filters, lang string) ([]entity.ProfessionalRole, int64, error)
}
