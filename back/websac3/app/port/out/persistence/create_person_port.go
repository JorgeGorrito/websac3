package persistence

import (
	"websac3/app/domain/entity"
)

type CreatePersonPort interface {
	Create(person *entity.Person, tx Transaction) error
}
