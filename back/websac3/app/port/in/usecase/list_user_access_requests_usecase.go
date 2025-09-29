package usecase

import (
	"websac3/app/domain/entity"
	commonfilter "websac3/common/filter"
	"websac3/common/paginator"
)

type ListUserAccessRequestsUseCase interface {
	Execute(userID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, lang string) ([]entity.AccessRequest, uint, error)
}
