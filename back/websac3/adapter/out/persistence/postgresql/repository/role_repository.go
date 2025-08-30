package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type RoleRepository struct {
	Repository
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (r *RoleRepository) GetByName(name string, ctx _db.Context) (entity.Role, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.Role{}, err
	}

	var role model.Role
	if err := dbCtx.DB().
		Model(&model.Role{}).
		Where("name = ?", name).
		First(&role).Error; err != nil {
		return entity.Role{}, err
	}

	roleEntity, err := mapper.Map[model.Role, entity.Role](&role)
	if err != nil {
		return entity.Role{}, err
	}

	return roleEntity, nil
}
