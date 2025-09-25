package repository

import (
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
	"websac3/common/mapper"
)

type HigherEducationInstitutionRepository struct {
	Repository
}

func NewHigherEducationInstitutionRepository() *HigherEducationInstitutionRepository {
	return &HigherEducationInstitutionRepository{}
}

func (r *HigherEducationInstitutionRepository) GetByFilters(
	page uint,
	perPage uint,
	filters filter.Filters,
	ctx db.Context) ([]entity.HigherEducationInstitution, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.HigherEducationInstitution{})
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

	var higherEducationInstitutions []model.HigherEducationInstitution
	if err := dbCtx.DB().
		Preload("Department").
		Preload("Municipality").
		Preload("Ownership").
		Preload("InstitutionalCategory").
		Find(&higherEducationInstitutions).Error; err != nil {
		return nil, 0, err
	}

	var results []entity.HigherEducationInstitution
	for _, institution := range higherEducationInstitutions {
		var higherEducationInstitution entity.HigherEducationInstitution
		if higherEducationInstitution, err = mapper.Map[model.HigherEducationInstitution, entity.HigherEducationInstitution](&institution); err != nil {
			return nil, 0, err
		}
		results = append(results, higherEducationInstitution)
	}

	return results, count, nil
}
