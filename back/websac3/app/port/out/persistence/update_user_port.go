package persistence

import "websac3/app/domain/entity"

type UpdateUserPort interface {
	UpdateByID(user *entity.User, userID uint, db Context) error
}
