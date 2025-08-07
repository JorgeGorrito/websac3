package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateAccessRequestPort interface {
	Create(accessRequest *entity.AccessRequest, db db.Context) error
}
