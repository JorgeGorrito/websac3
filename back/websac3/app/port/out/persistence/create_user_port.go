package persistence

import "websac3/app/domain/entity"

type CreateUserPort interface {
	Create(user *entity.User, tx Transaction) error
}
