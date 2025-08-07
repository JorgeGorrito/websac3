package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type UpdateUserPort interface {
	UpdateByID(user *entity.User, userID uint, db db.Context) error
}
