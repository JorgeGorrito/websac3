package usecase

import "websac3/app/domain/entity"

type LoginUseCase interface {
	Execute(user entity.User, lang string) (entity.User, error)
}
