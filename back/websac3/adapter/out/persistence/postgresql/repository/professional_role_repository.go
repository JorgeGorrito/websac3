package repository

import (
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
	"websac3/common/mapper"
)

type ProfessionalRoleRepository struct {
	Repository
}

func NewProfessionalRoleRepository() *ProfessionalRoleRepository {
	return &ProfessionalRoleRepository{}
}

func (r *ProfessionalRoleRepository) GetByID(id uint, lang string, ctx _db.Context) (entity.ProfessionalRole, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.ProfessionalRole{}, err
	}

	var professionalRole model.ProfessionalRole
	if err := dbCtx.DB().
		Preload("KnowledgeAreas.KnowledgeArea.Names").
		Preload("KnowledgeAreas.Topics.Topic.Names").
		Where("id = ?", id).
		First(&professionalRole).Error; err != nil {
		return entity.ProfessionalRole{}, err
	}

	// Use language-specific mapper if available, otherwise fallback to default
	var professionalRoleEntity entity.ProfessionalRole
	if langMapper := mapper.GetProfessionalRoleMapperWithLanguage(lang); langMapper != nil {
		professionalRoleEntity, err = langMapper(&professionalRole)
		if err != nil {
			return entity.ProfessionalRole{}, err
		}
	} else {
		// Fallback to default mapper
		if professionalRoleEntity, err = mapper.Map[model.ProfessionalRole, entity.ProfessionalRole](&professionalRole); err != nil {
			return entity.ProfessionalRole{}, err
		}
	}

	return professionalRoleEntity, nil
}

func (r *ProfessionalRoleRepository) GetByFilters(
	page uint,
	perPage uint,
	filters filter.Filters,
	ctx _db.Context,
) ([]entity.ProfessionalRole, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.ProfessionalRole{})
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
		Preload("KnowledgeAreas.KnowledgeArea.Names").
		Preload("KnowledgeAreas.Topics.Topic.Names").
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var professionalRoles []model.ProfessionalRole
	if err := dbCtx.DB().Find(&professionalRoles).Error; err != nil {
		return nil, 0, err
	}

	var results []entity.ProfessionalRole
	for _, role := range professionalRoles {
		var professionalRoleEntity entity.ProfessionalRole
		if professionalRoleEntity, err = mapper.Map[model.ProfessionalRole, entity.ProfessionalRole](&role); err != nil {
			return nil, 0, err
		}
		results = append(results, professionalRoleEntity)
	}

	return results, count, nil
}
