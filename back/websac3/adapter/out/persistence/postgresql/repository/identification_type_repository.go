package repository

import (
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
	"websac3/common/mapper"
)

type IdentificationTypeRepository struct {
	Repository
}

func NewIdentificationTypeRepository() *IdentificationTypeRepository {
	return &IdentificationTypeRepository{}
}

func (r *IdentificationTypeRepository) GetByFilters(
	page uint,
	perPage uint,
	filters filter.Filters,
	ctx _db.Context,
) ([]entity.IdentificationType, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.IdentificationType{})
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

	var identificationTypes []model.IdentificationType
	if err := dbCtx.DB().Find(&identificationTypes).Error; err != nil {
		return nil, 0, err
	}

	var results []entity.IdentificationType
	for _, idType := range identificationTypes {
		var identificationTypeEntity entity.IdentificationType
		if identificationTypeEntity, err = mapper.Map[model.IdentificationType, entity.IdentificationType](&idType); err != nil {
			return nil, 0, err
		}
		results = append(results, identificationTypeEntity)
	}

	return results, count, nil
}
