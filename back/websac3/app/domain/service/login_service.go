package service

import (
	"crypto/sha256"
	"encoding/hex"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type LoginService struct {
	msgProvider        message.Provider
	persistenceManager db.Manager
	getUserPort        persistence.GetUserPort
}

func NewLoginService(
	msgProvider message.Provider,
	persistenceManager db.Manager,
	getUserPort persistence.GetUserPort,
) *LoginService {
	return &LoginService{
		msgProvider:        msgProvider,
		persistenceManager: persistenceManager,
		getUserPort:        getUserPort,
	}
}

func (ls *LoginService) Execute(user entity.User, lang string) (entity.User, error) {
	var userFound entity.User
	var err error = ls.persistenceManager.ExecuteInTransaction(func(tx db.Context) error {
		var err error
		userFound, err = ls.getUserPort.GetByEmail(user.Email, tx)
		if err != nil {
			return err
		}

		pwdHash := sha256.Sum256([]byte(user.Password))
		pwdHashHex := hex.EncodeToString(pwdHash[:])

		if !userFound.IsRegistered() || !userFound.IsPasswordHashEqual(pwdHashHex) || !userFound.IsActive() {
			return errs.NewValidationError(
				ls.msgProvider.
					WithLang(lang).
					GetMessage("login", "invalid_user_password"),
			)
		}

		return nil
	})

	return userFound, err
}
