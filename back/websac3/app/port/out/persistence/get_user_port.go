package persistence

import "websac3/app/domain/entity"

type GetUserPort interface {
	GetByID(ID uint, db Context) (entity.User, error)
	GetByEmail(email string, db Context) (entity.User, error)
	GetByDNI(identificationType uint, identificationNumber string, db Context) (entity.User, error)
}
