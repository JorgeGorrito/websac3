package persistence

import (
	"websac3/app/domain/entity"
)

type GetStatusPort interface {
	GetByName(name string, tx Transaction) (entity.Status, error)
}
