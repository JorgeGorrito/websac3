package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	commonfilter "websac3/common/filter"
	"websac3/common/paginator"
)

type ListUserExpertConsultationsService struct {
	getExpertConsultationPort persistence.GetExpertConsultationPort
	persistenceManager        db.Manager
	messageProvider           message.Provider
}

func NewListUserExpertConsultationsService(
	getExpertConsultationPort persistence.GetExpertConsultationPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListUserExpertConsultationsService {
	return &ListUserExpertConsultationsService{
		getExpertConsultationPort: getExpertConsultationPort,
		persistenceManager:        persistenceManager,
		messageProvider:           messageProvider,
	}
}

func (s *ListUserExpertConsultationsService) Execute(userID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, lang string) ([]entity.ExpertConsultation, uint, error) {
	var expertConsultations []entity.ExpertConsultation = make([]entity.ExpertConsultation, 0)
	var total uint
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener solicitudes de asesoría del usuario autenticado
		var err error
		expertConsultations, total, err = s.getExpertConsultationPort.GetByUserID(userID, paginationParams, filters, ctx)
		if err != nil {
			expertConsultations = nil
			return err
		}

		// Verificar si no se encontraron solicitudes
		if len(expertConsultations) == 0 {
			err = errs.NewNotFoundError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("list_user_expert_consultations", "not_found"),
			)
		}

		return err
	})
	return expertConsultations, total, err
}
