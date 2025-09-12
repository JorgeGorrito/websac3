package repository

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"

	"gorm.io/gorm"
)

type ProfessionalRoleRepository struct{ Repository }

func NewProfessionalRoleRepository() *ProfessionalRoleRepository {
	return &ProfessionalRoleRepository{}
}

func (r *ProfessionalRoleRepository) GetByID(id uint, lang string, ctx _db.Context) (entity.ProfessionalRole, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.ProfessionalRole{}, err
	}

	var role model.ProfessionalRole
	if err := dbCtx.DB().
		Model(&model.ProfessionalRole{}).
		Preload("KnowledgeAreas").
		Preload("KnowledgeAreas.KnowledgeArea").
		Preload("KnowledgeAreas.KnowledgeArea.Names", "lang = ?", lang).
		Preload("KnowledgeAreas.Topics").
		Preload("KnowledgeAreas.Topics.Topic").
		Preload("KnowledgeAreas.Topics.Topic.Names", "lang = ?", lang).
		Preload("KnowledgeAreas.Topics.Topic.KnowledgeArea").
		Preload("KnowledgeAreas.Topics.Topic.KnowledgeArea.Names", "lang = ?", lang).
		Where("id = ?", id).
		First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ProfessionalRole{}, nil
		}
		return entity.ProfessionalRole{}, err
	}

	// Map to domain (basic fields; KnowledgeAreaExpected se arma en dominio si es necesario)
	return mapper.Map[model.ProfessionalRole, entity.ProfessionalRole](&role)
}
