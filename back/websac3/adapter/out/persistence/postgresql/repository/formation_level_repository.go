package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type FormationLevelRepository struct {
	Repository
}

func NewFormationLevelRepository() *FormationLevelRepository {
	return &FormationLevelRepository{}
}

func (r *FormationLevelRepository) GetAll(
	page uint,
	perPage uint,
	name string,
	ctx _db.Context,
) ([]entity.FormationLevel, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var formationLevels []model.FormationLevel
	var total int64

	// Build query
	query := dbCtx.DB().Model(&model.FormationLevel{})

	// Apply name filter if provided
	if name != "" {
		query = query.Joins("JOIN formation_level_names ON formation_levels.id = formation_level_names.formation_level_id").
			Where("formation_level_names.name ILIKE ?", "%"+name+"%").
			Group("formation_levels.id")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * perPage

	// Execute query with pagination
	if err := query.
		Preload("Names").
		Offset(int(offset)).
		Limit(int(perPage)).
		Order("formation_levels.id ASC").
		Find(&formationLevels).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities
	var formationLevelEntities []entity.FormationLevel
	for _, formationLevel := range formationLevels {
		formationLevelEntity, err := mapper.Map[model.FormationLevel, entity.FormationLevel](&formationLevel)
		if err != nil {
			return nil, 0, err
		}
		formationLevelEntities = append(formationLevelEntities, formationLevelEntity)
	}

	return formationLevelEntities, total, nil
}

func (r *FormationLevelRepository) GetAllWithLang(
	page uint,
	perPage uint,
	name string,
	lang string,
	ctx _db.Context,
) ([]entity.FormationLevel, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var formationLevels []model.FormationLevel
	var total int64

	// Build query
	query := dbCtx.DB().Model(&model.FormationLevel{})

	// Apply name filter if provided
	if name != "" {
		query = query.Joins("JOIN formation_level_names ON formation_levels.id = formation_level_names.formation_level_id").
			Where("formation_level_names.name ILIKE ?", "%"+name+"%").
			Group("formation_levels.id")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * perPage

	// Execute query with pagination
	if err := query.
		Preload("Names").
		Offset(int(offset)).
		Limit(int(perPage)).
		Order("formation_levels.id ASC").
		Find(&formationLevels).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities using language-specific mapper
	var formationLevelEntities []entity.FormationLevel
	for _, formationLevel := range formationLevels {
		formationLevelEntity, err := mapper.GetFormationLevelMapperWithLanguage(lang)(&formationLevel)
		if err != nil {
			return nil, 0, err
		}
		formationLevelEntities = append(formationLevelEntities, formationLevelEntity)
	}

	return formationLevelEntities, total, nil
}
