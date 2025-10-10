package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetExpertConsultationByIDService struct {
	getExpertConsultationPort persistence.GetExpertConsultationPort
	persistenceManager        db.Manager
	messageProvider           message.Provider
}

func NewGetExpertConsultationByIDService(
	getExpertConsultationPort persistence.GetExpertConsultationPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *GetExpertConsultationByIDService {
	return &GetExpertConsultationByIDService{
		getExpertConsultationPort: getExpertConsultationPort,
		persistenceManager:        persistenceManager,
		messageProvider:           messageProvider,
	}
}

func (s *GetExpertConsultationByIDService) Execute(consultationID uint, lang string) (entity.ExpertConsultation, error) {
	var expertConsultation entity.ExpertConsultation
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		var err error
		expertConsultation, err = s.getExpertConsultationPort.GetByID(consultationID, ctx)
		if err != nil {
			return errs.NewNotFoundError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("expert_consultation", "not_found"),
			)
		}
		return nil
	})
	return expertConsultation, err
}

