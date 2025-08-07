package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateUserPort interface {
	Create(user *entity.User, db db.Context) error
}
