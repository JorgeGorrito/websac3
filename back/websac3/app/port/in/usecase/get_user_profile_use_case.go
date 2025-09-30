package usecase

import "websac3/app/domain/entity"

type GetUserProfileUseCase interface {
	Execute(userID uint, lang string) (entity.User, error)
}
