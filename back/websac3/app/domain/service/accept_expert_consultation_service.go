package service

import (
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type AcceptExpertConsultationService struct {
	getExpertConsultationPort    persistence.GetExpertConsultationPort
	updateExpertConsultationPort persistence.UpdateExpertConsultationPort
	messageProvider              message.Provider
	persistenceManager           db.Manager
}

func NewAcceptExpertConsultationService(
	getExpertConsultationPort persistence.GetExpertConsultationPort,
	updateExpertConsultationPort persistence.UpdateExpertConsultationPort,
	messageProvider message.Provider,
	persistenceManager db.Manager,
) *AcceptExpertConsultationService {
	return &AcceptExpertConsultationService{
		getExpertConsultationPort:    getExpertConsultationPort,
		updateExpertConsultationPort: updateExpertConsultationPort,
		messageProvider:              messageProvider,
		persistenceManager:           persistenceManager,
	}
}

func (s *AcceptExpertConsultationService) Execute(consultationID uint, expertResponse string, expertID uint, lang string) error {
	return s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Verificar que la consulta existe
		consultation, err := s.getExpertConsultationPort.GetByID(consultationID, ctx)
		if err != nil {
			return err
		}

		// Verificar que la consulta está en estado pendiente
		if consultation.StatusID != 1 { // 1 = pending
			return errs.NewConflictError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("accept_expert_consultation", "not_pending"),
			)
		}

		// Verificar que la consulta no tiene ya un experto asignado
		if consultation.ExpertID != nil {
			return errs.NewConflictError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("accept_expert_consultation", "already_assigned"),
			)
		}

		// Aceptar la consulta
		err = s.updateExpertConsultationPort.AcceptConsultation(consultationID, expertResponse, expertID, ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
