package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ValidateEmailService struct {
	getAccessRequestPort    persistence.GetAccessRequestPort
	updateAccessRequestPort persistence.UpdateAccessRequestPort
	msgProvider             message.Provider
	persistenceManager      db.Manager
}

func NewValidateEmailService(
	getAccessRequestPort persistence.GetAccessRequestPort,
	updateAccessRequestPort persistence.UpdateAccessRequestPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) *ValidateEmailService {
	return &ValidateEmailService{
		getAccessRequestPort:    getAccessRequestPort,
		updateAccessRequestPort: updateAccessRequestPort,
		msgProvider:             msgProvider,
		persistenceManager:      persistenceManager,
	}
}

func (v *ValidateEmailService) Execute(validationToken string, lang string) error {
	return v.persistenceManager.ExecuteInTransaction(
		func(tx db.Context) error {
			var err error
			var accessRequestFound entity.AccessRequest

			if accessRequestFound, err = v.getAccessRequestPort.GetUnvalidatedEmailByToken(validationToken, tx); err != nil {
				return err
			}

			if !accessRequestFound.IsRegistered() {
				return errs.NewNotFoundError(
					v.msgProvider.
						WithLang(lang).
						GetMessage("validate_email", "email_not_found"),
				)
			}

			accessRequestFound.IsVerified = true
			if err := v.updateAccessRequestPort.Update(&accessRequestFound, tx); err != nil {
				return err
			}

			return nil
		},
	)
}
