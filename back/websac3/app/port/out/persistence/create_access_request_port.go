package persistence

import (
	"websac3/app/domain/entity"
)

type CreateAccessRequestPort interface {
	Create(accessRequest *entity.AccessRequest, tx Transaction) error
}
