package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ActivateUserService struct {
	getUserPort        persistence.GetUserPort
	updateUserPort     persistence.UpdateUserPort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewActivateUserService(
	getUserPort persistence.GetUserPort,
	updateUserPort persistence.UpdateUserPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ActivateUserService {
	return &ActivateUserService{
		getUserPort:        getUserPort,
		updateUserPort:     updateUserPort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *ActivateUserService) Execute(userID uint, lang string) error {
	var user entity.User
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		var err error

		// Verificar que el usuario existe
		user, err = s.getUserPort.GetByID(userID, ctx)
		if err != nil {
			return errs.NewNotFoundError(
				s.msgProvider.WithLang(lang).GetMessage("activate_user", "user_not_found"),
			)
		}

		// Verificar que el usuario esté desactivado
		if user.IsActive() {
			return errs.NewConflictError(
				s.msgProvider.WithLang(lang).GetMessage("activate_user", "user_already_active"),
			)
		}

		// Activar el usuario (establecer DeactivatedAt como nil)
		user.DeactivatedAt = nil

		// Actualizar en la base de datos
		err = s.updateUserPort.UpdateByID(&user, userID, ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
