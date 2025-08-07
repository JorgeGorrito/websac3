package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetUserPort interface {
	GetByID(ID uint, db db.Context) (entity.User, error)
	GetByEmail(email string, db db.Context) (entity.User, error)
	GetByDNI(identificationType uint, identificationNumber string, db db.Context) (entity.User, error)
}
