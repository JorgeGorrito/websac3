package persistence

import "websac3/app/domain/entity"

type UpdatePersonPort interface {
	UpdateById(person *entity.Person, personID uint, tx Transaction) error
}
