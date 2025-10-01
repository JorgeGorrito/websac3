package service

import (
	"crypto/sha256"
	"encoding/hex"
	"websac3/app/domain/errs"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ChangePasswordService struct {
	getUserPort        persistence.GetUserPort
	updateUserPort     persistence.UpdateUserPort
	msgProvider        message.Provider
	persistenceManager db.Manager
}

func NewChangePasswordService(
	getUserPort persistence.GetUserPort,
	updateUserPort persistence.UpdateUserPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) usecase.ChangePasswordUseCase {
	return &ChangePasswordService{
		getUserPort:        getUserPort,
		updateUserPort:     updateUserPort,
		msgProvider:        msgProvider,
		persistenceManager: persistenceManager,
	}
}

func (s *ChangePasswordService) Execute(cmd command.ChangePasswordCommand, lang string) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			// 1. Obtener el usuario actual
			user, err := s.getUserPort.GetByID(cmd.UserID, ctx)
			if err != nil {
				return err
			}

			if !user.IsRegistered() {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("change_password", "user_not_found"),
				)
			}

			// 2. Validar que el usuario esté activo
			if !user.IsActive() {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("change_password", "user_not_active"),
				)
			}

			// 3. Validar que la contraseña actual sea correcta
			currentPwdHash := sha256.Sum256([]byte(cmd.CurrentPassword))
			currentPwdHashHex := hex.EncodeToString(currentPwdHash[:])

			if !user.IsPasswordHashEqual(currentPwdHashHex) {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("change_password", "current_password_incorrect"),
				)
			}

			// 4. Validar que la nueva contraseña y su confirmación coincidan
			if cmd.NewPassword != cmd.ConfirmNewPassword {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("change_password", "passwords_mismatch"),
				)
			}

			// 5. Validar que la nueva contraseña sea diferente de la actual
			if cmd.CurrentPassword == cmd.NewPassword {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("change_password", "new_password_same_as_current"),
				)
			}

			// 6. Generar el hash de la nueva contraseña
			newPwdHash := sha256.Sum256([]byte(cmd.NewPassword))
			newPwdHashHex := hex.EncodeToString(newPwdHash[:])

			// 7. Actualizar la contraseña del usuario
			user.PasswordHash = newPwdHashHex
			if err := s.updateUserPort.UpdateByID(&user, cmd.UserID, ctx); err != nil {
				return err
			}

			return nil
		},
	)
}
