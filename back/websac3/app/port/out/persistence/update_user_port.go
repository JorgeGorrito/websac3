package persistence

import "websac3/app/domain/entity"

type UpdateUserPort interface {
	UpdateById(user *entity.User, userID uint, tx Transaction) error
}
