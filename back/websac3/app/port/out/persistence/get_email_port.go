package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetEmailPort interface {
	GetChunkNotSent(chunkSize uint, page uint, db db.Context) ([]entity.EmailNotification, error)
}
