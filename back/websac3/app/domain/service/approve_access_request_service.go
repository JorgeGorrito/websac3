package service

import (
	"errors"
	"time"
	"websac3/app/domain/constants"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/domain/notification/template/context"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
)

type ApproveAccessRequestService struct {
	statusEnum              enum.StatusEnum
	getUserPort             persistence.GetUserPort
	getAccessRequestPort    persistence.GetAccessRequestPort
	updateAccessRequestPort persistence.UpdateAccessRequestPort
	msgProvider             message.Provider
	templateProvider        template.Provider
	sendNotificationPort    notification.SendMailPort
	persistenceManager      db.Manager
}

func NewApproveAccessRequestService(
	statusEnum enum.StatusEnum,
	getUserPort persistence.GetUserPort,
	getAccessRequestPort persistence.GetAccessRequestPort,
	updateAccessRequestPort persistence.UpdateAccessRequestPort,
	msgProvider message.Provider,
	notificationPort notification.SendMailPort,
	persistenceManager db.Manager,
	templateProvider template.Provider,
) *ApproveAccessRequestService {
	return &ApproveAccessRequestService{
		statusEnum:              statusEnum,
		getUserPort:             getUserPort,
		getAccessRequestPort:    getAccessRequestPort,
		updateAccessRequestPort: updateAccessRequestPort,
		msgProvider:             msgProvider,
		sendNotificationPort:    notificationPort,
		persistenceManager:      persistenceManager,
		templateProvider:        templateProvider,
	}
}

func (s *ApproveAccessRequestService) Execute(requestTpApproveID uint, lang string) error {
	var err error = s.persistenceManager.ExecuteInTransaction(
		func(tx db.Context) error {
			var err error
			var accessRequest entity.AccessRequest
			var user entity.User

			accessRequest, err = s.getAccessRequestPort.GetByID(requestTpApproveID, tx)
			if err != nil {
				return err
			}

			user, err = s.getUserPort.GetByDNI(accessRequest.Applicant.IdentificationTypeID, accessRequest.Applicant.IdentificationNumber, tx)
			if err != nil {
				return err
			}
			if user.IsRegistered() {
				return errs.NewConflictError(s.msgProvider.WithLang(lang).GetMessage("approve_access_request", "user_alredy_registered"))
			}

			if !accessRequest.IsRegistered() || !accessRequest.HasEmailVerified() {
				return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("approve_access_request", "not_found"))
			}

			if accessRequest.IsApproved() {
				return errs.NewConflictError(s.msgProvider.WithLang(lang).GetMessage("approve_access_request", "already_approved"))
			}

			statusApproved, err := s.statusEnum.GetByName(constants.Approved)
			if err != nil {
				return err
			}
			accessRequest.StatusID = statusApproved.ID
			accessRequest.Status = &statusApproved

			if err = s.updateAccessRequestPort.Update(&accessRequest, tx); err != nil {
				return err
			}

			var templateApproved template.Template
			if templateApproved = s.templateProvider.GetByName("access_request_approved"); templateApproved == nil {
				return errors.New(s.msgProvider.WithLang(lang).GetMessage("template", "template_not_found", "access_request_approved"))
			}

			var templateToSend string
			var createUserURL string = accessRequest.CreateUserURL + accessRequest.CreateUserCode
			if templateToSend, err = templateApproved.Render(
				&context.AccessRequestApproved{
					NombrePersona:        accessRequest.Applicant.Name,
					EnlaceCreacionCuenta: createUserURL,
				},
			); err != nil {
				return err
			}

			accessRequest.EmailApproved = &entity.EmailNotification{
				To:        accessRequest.EmailValidation.To,
				Content:   templateToSend,
				Subject:   "Solicitud de acceso aprobada",
				CreatedAt: time.Now(),
			}
			if err = s.sendNotificationPort.Send(
				accessRequest.EmailApproved,
				tx,
			); err != nil {
				return err
			}

			accessRequest.EmailApprovedID = func() *uint {
				if accessRequest.EmailApproved == nil {
					return nil
				}
				id := accessRequest.EmailApproved.ID
				return &id
			}()
			if err = s.updateAccessRequestPort.Update(&accessRequest, tx); err != nil {
				return err
			}

			return nil
		},
	)
	return err
}
