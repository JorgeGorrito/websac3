package service

import (
	"time"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type DeactivateUserService struct {
	getUserPort        persistence.GetUserPort
	updateUserPort     persistence.UpdateUserPort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewDeactivateUserService(
	getUserPort persistence.GetUserPort,
	updateUserPort persistence.UpdateUserPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *DeactivateUserService {
	return &DeactivateUserService{
		getUserPort:        getUserPort,
		updateUserPort:     updateUserPort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *DeactivateUserService) Execute(userID uint, lang string) error {
	var user entity.User
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		var err error

		// Verificar que el usuario existe
		user, err = s.getUserPort.GetByID(userID, ctx)
		if err != nil {
			return errs.NewNotFoundError(
				s.msgProvider.WithLang(lang).GetMessage("deactivate_user", "user_not_found"),
			)
		}

		// Verificar que el usuario no esté ya desactivado
		if !user.IsActive() {
			return errs.NewConflictError(
				s.msgProvider.WithLang(lang).GetMessage("deactivate_user", "user_already_deactivated"),
			)
		}

		// Desactivar el usuario
		now := time.Now()
		user.DeactivatedAt = &now

		// Actualizar en la base de datos
		err = s.updateUserPort.UpdateByID(&user, userID, ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
