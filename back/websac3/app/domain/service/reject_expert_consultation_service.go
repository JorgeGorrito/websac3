package service

import (
	"time"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/domain/notification/template/context"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type RejectExpertConsultationService struct {
	getExpertConsultationPort    persistence.GetExpertConsultationPort
	updateExpertConsultationPort persistence.UpdateExpertConsultationPort
	messageProvider              message.Provider
	persistenceManager           db.Manager
	templateProvider             template.Provider
	sendNotificationPort         notification.SendMailPort
}

func NewRejectExpertConsultationService(
	getExpertConsultationPort persistence.GetExpertConsultationPort,
	updateExpertConsultationPort persistence.UpdateExpertConsultationPort,
	messageProvider message.Provider,
	persistenceManager db.Manager,
	templateProvider template.Provider,
	sendNotificationPort notification.SendMailPort,
) *RejectExpertConsultationService {
	return &RejectExpertConsultationService{
		getExpertConsultationPort:    getExpertConsultationPort,
		updateExpertConsultationPort: updateExpertConsultationPort,
		messageProvider:              messageProvider,
		persistenceManager:           persistenceManager,
		templateProvider:             templateProvider,
		sendNotificationPort:         sendNotificationPort,
	}
}

func (s *RejectExpertConsultationService) Execute(consultationID uint, expertResponse string, expertID uint, lang string) error {
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
					GetMessage("reject_expert_consultation", "not_pending"),
			)
		}

		// Verificar que la consulta no tiene ya un experto asignado
		if consultation.ExpertID != nil {
			return errs.NewConflictError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("reject_expert_consultation", "already_assigned"),
			)
		}

		// Rechazar la consulta
		err = s.updateExpertConsultationPort.RejectConsultation(consultationID, expertResponse, expertID, ctx)
		if err != nil {
			return err
		}

		// Enviar notificación por correo
		var templateNotification template.Template
		if templateNotification = s.templateProvider.GetByNameAndLang("expert_consultation_notification", lang); templateNotification == nil {
			// No retornamos error para no afectar la operación principal
			return nil
		}

		// Crear el contexto para la plantilla
		templateContext := context.ExpertConsultationNotificationContext{
			RequesterName:                 consultation.Requester.Person.Name + " " + consultation.Requester.Person.Lastname,
			RequesterEmail:                consultation.Requester.Email,
			RequesterInstitutionName:      consultation.Requester.Person.HigherEducationInstitution.Name,
			RequesterInstitutionOwnership: consultation.Requester.Person.HigherEducationInstitution.Ownership.Name,
			RequesterJobPosition:          consultation.Requester.Person.JobPosition,
			DegreeProgramName:             consultation.DegreeProgram.Name,
			DegreeProgramSnies:            consultation.DegreeProgram.Snies,
			ReportScore:                   consultation.Report.Score,
			RequestMessage:                *consultation.RequestMessage,
			ExpertResponse:                expertResponse,
			Status:                        "rejected",
			CreatedAt:                     consultation.CreatedAt.Format("2006-01-02 15:04:05"),
			AnsweredAt:                    time.Now().Format("2006-01-02 15:04:05"),
		}

		var templateToSend string
		if templateToSend, err = templateNotification.Render(templateContext); err != nil {
			// No retornamos error para no afectar la operación principal
			return nil
		}

		// Determinar el asunto del correo según el idioma
		var subject string
		if lang == "es" {
			subject = "Solicitud de Asesoría Rechazada - WEBSAC3"
		} else {
			subject = "Expert Consultation Request Rejected - WEBSAC3"
		}

		// Crear notificación de email
		emailNotification := &entity.EmailNotification{
			To:          consultation.Requester.Email,
			Subject:     subject,
			Content:     templateToSend,
			Attachments: []entity.EmailAttachment{}, // Sin adjuntos por ahora
		}

		// Enviar notificación
		if err = s.sendNotificationPort.Send(emailNotification, ctx); err != nil {
			// No retornamos error para no afectar la operación principal
		}

		return nil
	})
}
