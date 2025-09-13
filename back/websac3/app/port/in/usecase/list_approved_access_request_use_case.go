package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/filter"
)

type ListApprovedAccessRequestUseCase interface {
	Execute(page, perPage uint, filters filter.Filters, userID uint, permissions []string, lang string) ([]entity.AccessRequest, int64, error)
}
