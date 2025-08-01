package usecase

import (
	"websac3/app/domain/entity"
)

type CreateAccessRequestUseCase interface {
	Execute(request entity.AccessRequest, redirectURLTo string, lang string) error
}
