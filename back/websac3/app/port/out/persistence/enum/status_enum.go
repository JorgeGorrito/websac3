package enum

import "websac3/app/domain/entity"

type StatusEnum interface {
	GetByName(name string) (entity.Status, error)
}
