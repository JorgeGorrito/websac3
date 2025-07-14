package persistence

import "websac3/app/domain/entity"

type CreateUserPort interface {
	Create(user *entity.User, db Context) error
}
