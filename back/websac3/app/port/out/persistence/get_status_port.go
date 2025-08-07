package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetStatusPort interface {
	GetByName(name string, db db.Context) (entity.Status, error)
}
