package persistence

import "websac3/app/domain/entity"

type GetUserPort interface {
	GetByEmail(email string, tx Transaction) (entity.User, error)
}
