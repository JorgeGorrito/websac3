package usecase

import (
	"websac3/app/domain/entity"
)

type ListFormationLevelUseCase interface {
	Execute(page, perPage uint, name string, lang string) ([]entity.FormationLevel, int64, error)
}
