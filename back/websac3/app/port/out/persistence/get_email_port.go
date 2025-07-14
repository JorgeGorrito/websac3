package persistence

import "websac3/app/domain/entity"

type GetEmailPort interface {
	GetChunkNotSent(chunkSize uint, page uint, db Context) ([]entity.EmailNotification, error)
}
