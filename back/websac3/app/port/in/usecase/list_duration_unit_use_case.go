package usecase

import (
	"websac3/app/domain/entity"
)

type ListDurationUnitUseCase interface {
	Execute(page, perPage uint, name string, lang string) ([]entity.DurationUnit, int64, error)
}
