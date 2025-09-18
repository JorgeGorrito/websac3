package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/common/paginator"
)

type ListUsersUseCase interface {
	Execute(query query.ListUsersQuery, lang string) (*paginator.Page[entity.User], error)
}
