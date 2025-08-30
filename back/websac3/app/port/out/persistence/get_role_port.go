package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetRolePort interface {
	GetByName(name string, ctx db.Context) (entity.Role, error)
}
