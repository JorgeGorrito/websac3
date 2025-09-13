package repository

import (
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
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

func (r *RoleRepository) GetByID(id uint, ctx _db.Context) (entity.Role, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.Role{}, err
	}

	var role model.Role
	if err := dbCtx.DB().
		Model(&model.Role{}).
		Where("id = ?", id).
		First(&role).Error; err != nil {
		return entity.Role{}, err
	}

	roleEntity, err := mapper.Map[model.Role, entity.Role](&role)
	if err != nil {
		return entity.Role{}, err
	}

	return roleEntity, nil
}

func (r *RoleRepository) GetByFilters(
	page uint,
	perPage uint,
	filters filter.Filters,
	ctx _db.Context,
) ([]entity.Role, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.Role{})
	dbCtx.DBSet(baseQuery)

	dbCtxFiltered, err := filters.Apply(dbCtx, psqlfilter.FiltersRegistry)
	if err != nil {
		return nil, 0, err
	}

	dbCtx, err = r.CastDbContext(dbCtxFiltered)
	if err != nil {
		return nil, 0, err
	}

	var count int64
	if err := dbCtx.DB().Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var roles []model.Role
	if err := dbCtx.DB().Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	var results []entity.Role
	for _, role := range roles {
		var roleEntity entity.Role
		if roleEntity, err = mapper.Map[model.Role, entity.Role](&role); err != nil {
			return nil, 0, err
		}
		results = append(results, roleEntity)
	}

	return results, count, nil
}
