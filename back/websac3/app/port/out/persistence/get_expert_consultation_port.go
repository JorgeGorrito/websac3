package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	commonfilter "websac3/common/filter"
	"websac3/common/paginator"
)

type GetExpertConsultationPort interface {
	GetByID(id uint, ctx db.Context) (entity.ExpertConsultation, error)
	GetByUserID(userID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, ctx db.Context) ([]entity.ExpertConsultation, uint, error)
	GetPendingConsultations(paginationParams paginator.PaginationParams, filters commonfilter.Params, ctx db.Context) ([]entity.ExpertConsultation, uint, error)
	GetAnsweredConsultationsByExpertID(expertID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, ctx db.Context) ([]entity.ExpertConsultation, uint, error)
}

type UpdateExpertConsultationPort interface {
	AcceptConsultation(consultationID uint, expertResponse string, expertID uint, ctx db.Context) error
	RejectConsultation(consultationID uint, expertResponse string, expertID uint, ctx db.Context) error
}
