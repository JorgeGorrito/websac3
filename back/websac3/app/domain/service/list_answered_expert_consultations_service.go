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

type ListAnsweredExpertConsultationsService struct {
	getExpertConsultationPort persistence.GetExpertConsultationPort
	persistenceManager        db.Manager
	messageProvider           message.Provider
}

func NewListAnsweredExpertConsultationsService(
	getExpertConsultationPort persistence.GetExpertConsultationPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListAnsweredExpertConsultationsService {
	return &ListAnsweredExpertConsultationsService{
		getExpertConsultationPort: getExpertConsultationPort,
		persistenceManager:        persistenceManager,
		messageProvider:           messageProvider,
	}
}

func (s *ListAnsweredExpertConsultationsService) Execute(userID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, lang string) ([]entity.ExpertConsultation, uint, error) {
	var expertConsultations []entity.ExpertConsultation = make([]entity.ExpertConsultation, 0)
	var total uint
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener solicitudes de asesoría respondidas por el experto
		var err error
		expertConsultations, total, err = s.getExpertConsultationPort.GetAnsweredConsultationsByExpertID(userID, paginationParams, filters, ctx)
		if err != nil {
			expertConsultations = nil
			return err
		}

		// Verificar si no se encontraron solicitudes
		if len(expertConsultations) == 0 {
			err = errs.NewNotFoundError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("list_answered_expert_consultations", "not_found"),
			)
		}

		return err
	})
	return expertConsultations, total, err
}


