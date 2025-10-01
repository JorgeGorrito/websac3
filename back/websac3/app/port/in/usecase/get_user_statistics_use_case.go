package usecase

import "websac3/app/domain/entity"

type GetUserStatisticsUseCase interface {
	Execute(userID uint, lang string) (entity.UserStatistics, error)
}
