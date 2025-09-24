package usecase

import (
	"websac3/app/domain/entity"
)

type CreateExpertConsultationUseCase interface {
	Execute(request entity.ExpertConsultation, lang string) error
}
