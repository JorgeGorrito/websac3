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
	roleEnum                enum.RoleEnum
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
	roleEnum enum.RoleEnum,
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
		roleEnum:                roleEnum,
	}
}

func (s *ApproveAccessRequestService) Execute(requestTpApproveID uint, roleID uint, lang string) error {
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

			// Validar que el roleID sea válido
			if roleID == 0 {
				return errs.NewValidationError(s.msgProvider.WithLang(lang).GetMessage("approve_access_request", "role_required"))
			}

			// Obtener el rol por ID
			roleEntity, err := s.roleEnum.GetByID(roleID)
			if err != nil {
				return errs.NewValidationError(s.msgProvider.WithLang(lang).GetMessage("approve_access_request", "invalid_role"))
			}

			statusApproved, err := s.statusEnum.GetByName(constants.Approved)
			if err != nil {
				return err
			}
			accessRequest.StatusID = statusApproved.ID
			accessRequest.Status = &statusApproved

			// Asignar el rol aprobado
			accessRequest.ApprovedRoleID = &roleEntity.ID
			accessRequest.ApprovedRole = &roleEntity

			if err = s.updateAccessRequestPort.Update(&accessRequest, tx); err != nil {
				return err
			}

			var templateApproved template.Template
			if templateApproved = s.templateProvider.GetByNameAndLang("access_request_approved", accessRequest.Lang); templateApproved == nil {
				return errors.New(s.msgProvider.WithLang(accessRequest.Lang).GetMessage("template", "template_not_found", "access_request_approved"))
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
				Subject:   s.msgProvider.WithLang(accessRequest.Lang).GetMessage("approve_access_request", "approve_success"),
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
