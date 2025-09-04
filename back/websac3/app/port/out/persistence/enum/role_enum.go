package enum

import "websac3/app/domain/entity"

type RoleEnum interface {
	GetByName(name string) (entity.Role, error)
	GetByID(id uint) (entity.Role, error)
}
