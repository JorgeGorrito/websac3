package usecase

import "websac3/app/domain/entity"

type GetExpertConsultationByIDUseCase interface {
	Execute(consultationID uint, lang string) (entity.ExpertConsultation, error)
}

