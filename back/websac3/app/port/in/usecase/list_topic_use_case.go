package usecase

import (
	"websac3/app/domain/entity"
)

type ListTopicUseCase interface {
	Execute(page, perPage uint, name string, lang string) ([]entity.Topic, int64, error)
}
