package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetUserByIDService struct {
	getUserPort        persistence.GetUserPort
	msgProvider        message.Provider
	persistenceManager db.Manager
}

func NewGetUserByIDService(
	getUserPort persistence.GetUserPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) *GetUserByIDService {
	return &GetUserByIDService{
		getUserPort:        getUserPort,
		msgProvider:        msgProvider,
		persistenceManager: persistenceManager,
	}
}

func (s *GetUserByIDService) Execute(ID uint, lang string) (entity.User, error) {
	var user entity.User
	var err error = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			var err error
			user, err = s.getUserPort.GetByID(ID, dbCtx)
			if err != nil {
				user = entity.User{}
				return err
			}

			if !user.IsRegistered() {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("user", "user_not_found"),
				)
			}

			return nil
		},
	)
	return user, err
}
