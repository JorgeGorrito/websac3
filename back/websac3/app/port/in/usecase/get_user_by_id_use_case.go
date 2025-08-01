package usecase

import "websac3/app/domain/entity"

type GetUserByIDUseCase interface {
	Execute(ID uint, lang string) (entity.User, error)
}
