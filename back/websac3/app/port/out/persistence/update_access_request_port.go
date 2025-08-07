package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type UpdateAccessRequestPort interface {
	Update(request *entity.AccessRequest, ctx db.Context) error
}
