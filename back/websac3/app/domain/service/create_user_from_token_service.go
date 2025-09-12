package service

import (
	"crypto/sha256"
	"encoding/hex"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
)

type CreateUserFromTokenService struct {
	msgProvider             message.Provider
	createUserPort          persistence.CreateUserPort
	createPersonPort        persistence.CreatePersonPort
	getAccessRequestPort    persistence.GetAccessRequestPort
	updateAccessRequestPort persistence.UpdateAccessRequestPort
	getUserPort             persistence.GetUserPort
	roleEnum                enum.RoleEnum
	persistenceManager      db.Manager
}

func NewCreateUserFromTokenService(
	getAccessRequestPort persistence.GetAccessRequestPort,
	createUserPort persistence.CreateUserPort,
	createPersonPort persistence.CreatePersonPort,
	updateAccessRequestPort persistence.UpdateAccessRequestPort,
	getUserPort persistence.GetUserPort,
	roleEnum enum.RoleEnum,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) *CreateUserFromTokenService {
	return &CreateUserFromTokenService{
		msgProvider:             msgProvider,
		createUserPort:          createUserPort,
		createPersonPort:        createPersonPort,
		getAccessRequestPort:    getAccessRequestPort,
		updateAccessRequestPort: updateAccessRequestPort,
		getUserPort:             getUserPort,
		roleEnum:                roleEnum,
		persistenceManager:      persistenceManager,
	}
}

func (s *CreateUserFromTokenService) Execute(token string, password string, lang string) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			// 1. Obtener la solicitud de acceso usando el token
			accessRequest, err := s.getAccessRequestPort.GetByCreateUserToken(token, ctx)
			if err != nil {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("create_user", "access_request_not_found"),
				)
			}

			// 2. Verificar que la solicitud esté aprobada
			if !accessRequest.IsApproved() {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("create_user", "access_request_not_approved"),
				)
			}

			// 3. Verificar que el email esté verificado
			if !accessRequest.HasEmailVerified() {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("create_user", "email_not_verified"),
				)
			}

			// 4. Verificar que no exista ya un usuario con el email de la solicitud
			if accessRequest.EmailValidation != nil {
				existingUser, err := s.getUserPort.GetByEmail(accessRequest.EmailValidation.To, ctx)
				if err == nil && existingUser.ID != 0 {
					// Si no hay error y el usuario existe, significa que ya se creó un usuario para este email
					return errs.NewConflictError(
						s.msgProvider.
							WithLang(lang).
							GetMessage("create_user", "user_already_exists"),
					)
				}
				// Si hay error, asumimos que el usuario no existe (esto es normal)
			}

			// 5. Verificar que el email esté disponible
			if accessRequest.EmailValidation == nil {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("create_user", "email_not_verified"),
				)
			}

			// 6. Obtener el rol aprobado en la solicitud
			if accessRequest.ApprovedRole == nil {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("create_user", "role_not_defined"),
				)
			}

			userRole := *accessRequest.ApprovedRole

			// 7. Crear la persona si no existe
			if accessRequest.Applicant != nil && !accessRequest.Applicant.IsRegistered() {
				if err := s.createPersonPort.Create(accessRequest.Applicant, ctx); err != nil {
					return err
				}
			}

			// 8. Generar el hash de la contraseña
			passwordHash := sha256.Sum256([]byte(password))
			passwordHashHex := hex.EncodeToString(passwordHash[:])

			// 9. Crear el usuario
			user := &entity.User{
				Password:     password,
				PasswordHash: passwordHashHex,
				Email:        accessRequest.EmailValidation.To,
				RoleID:       userRole.ID,
				Role:         &userRole,
				PersonID:     accessRequest.ApplicantID,
				Person:       accessRequest.Applicant,
			}

			if err := s.createUserPort.Create(user, ctx); err != nil {
				return err
			}

			return nil
		},
	)
}
