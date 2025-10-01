package service

import (
	"errors"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type DeleteDegreeProgramService struct {
	persistenceManager      db.Manager
	deleteDegreeProgramPort persistence.DeleteDegreeProgramPort
	getDegreeProgramPort    persistence.GetDegreeProgramPort
	getUserPort             persistence.GetUserPort
	msgProvider             message.Provider
}

func NewDeleteDegreeProgramService(
	persistenceManager db.Manager,
	deleteDegreeProgramPort persistence.DeleteDegreeProgramPort,
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	getUserPort persistence.GetUserPort,
	msgProvider message.Provider,
) *DeleteDegreeProgramService {
	return &DeleteDegreeProgramService{
		persistenceManager:      persistenceManager,
		deleteDegreeProgramPort: deleteDegreeProgramPort,
		getDegreeProgramPort:    getDegreeProgramPort,
		getUserPort:             getUserPort,
		msgProvider:             msgProvider,
	}
}

func (s *DeleteDegreeProgramService) Execute(
	degreeProgramID uint,
	userID uint,
	lang string,
) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			// Get the user to check if they are admin
			user, err := s.getUserPort.GetByID(userID, ctx)
			if err != nil {
				return err
			}

			// Get the degree program to check ownership and permissions
			degreeProgram, err := s.getDegreeProgramPort.GetByIDWithLang(degreeProgramID, lang, ctx)
			if err != nil {
				if errors.Is(err, errs.NotFoundError) {
					return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("delete_degree_program", "degree_program_not_found"))
				}
				return err
			}

			// Check if the degree program is already deleted
			if degreeProgram.DeletedAt != nil {
				return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("delete_degree_program", "degree_program_not_found"))
			}

			// Use business rules from entities to check permissions
			if !user.IsAdmin() && !degreeProgram.CanBeDeletedBy(userID) {
				return errs.NewValidationError(s.msgProvider.WithLang(lang).GetMessage("delete_degree_program", "insufficient_permissions"))
			}

			// Proceed with deletion
			if err := s.deleteDegreeProgramPort.DeleteByID(degreeProgramID, ctx); err != nil {
				return err
			}
			return nil
		},
	)
}
