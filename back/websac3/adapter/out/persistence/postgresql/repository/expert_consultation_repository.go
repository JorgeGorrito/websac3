package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type ExpertConsultationRepository struct {
	Repository
}

func NewExpertConsultationRepository() *ExpertConsultationRepository {
	return &ExpertConsultationRepository{}
}

func (e *ExpertConsultationRepository) Create(expertConsultationToSave *entity.ExpertConsultation, ctx _db.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var expertConsultation model.ExpertConsultation
	if expertConsultation, err = mapper.Map[entity.ExpertConsultation, model.ExpertConsultation](expertConsultationToSave); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&expertConsultation).
		Error; err != nil {
		return err
	}
	expertConsultationToSave.ID = expertConsultation.ID

	return nil
}

func (e *ExpertConsultationRepository) GetByID(id uint, ctx _db.Context) (entity.ExpertConsultation, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return entity.ExpertConsultation{}, err
	}

	var expertConsultation model.ExpertConsultation
	if err := dbCtx.DB().
		Preload("Requester").
		Preload("Requester.Person").
		Preload("DegreeProgram").
		Preload("Report").
		Preload("Expert").
		Preload("Expert.Person").
		Preload("Status").
		Preload("Status.Names").
		First(&expertConsultation, id).
		Error; err != nil {
		return entity.ExpertConsultation{}, err
	}

	return mapper.Map[model.ExpertConsultation, entity.ExpertConsultation](&expertConsultation)
}

func (e *ExpertConsultationRepository) GetByRequesterID(requesterID uint, ctx _db.Context) ([]entity.ExpertConsultation, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var expertConsultations []model.ExpertConsultation
	if err := dbCtx.DB().
		Preload("Requester").
		Preload("Requester.Person").
		Preload("DegreeProgram").
		Preload("Report").
		Preload("Expert").
		Preload("Expert.Person").
		Preload("Status").
		Preload("Status.Names").
		Where("requester_id = ?", requesterID).
		Order("created_at DESC").
		Find(&expertConsultations).
		Error; err != nil {
		return nil, err
	}

	return mapper.Map[[]model.ExpertConsultation, []entity.ExpertConsultation](&expertConsultations)
}

func (e *ExpertConsultationRepository) GetByStatus(status string, ctx _db.Context) ([]entity.ExpertConsultation, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var expertConsultations []model.ExpertConsultation
	if err := dbCtx.DB().
		Preload("Requester").
		Preload("Requester.Person").
		Preload("DegreeProgram").
		Preload("Report").
		Preload("Expert").
		Preload("Expert.Person").
		Preload("Status").
		Preload("Status.Names").
		Joins("JOIN expert_consultation_statuses ON expert_consultations.status_id = expert_consultation_statuses.id").
		Joins("JOIN expert_consultation_status_names ON expert_consultation_statuses.id = expert_consultation_status_names.expert_consultation_status_id").
		Where("expert_consultation_status_names.name = ?", status).
		Order("created_at DESC").
		Find(&expertConsultations).
		Error; err != nil {
		return nil, err
	}

	return mapper.Map[[]model.ExpertConsultation, []entity.ExpertConsultation](&expertConsultations)
}

func (e *ExpertConsultationRepository) Update(expertConsultationToUpdate *entity.ExpertConsultation, ctx _db.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var expertConsultation model.ExpertConsultation
	if expertConsultation, err = mapper.Map[entity.ExpertConsultation, model.ExpertConsultation](expertConsultationToUpdate); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Save(&expertConsultation).
		Error; err != nil {
		return err
	}

	return nil
}

func (e *ExpertConsultationRepository) Delete(id uint, ctx _db.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	if err := dbCtx.DB().
		Delete(&model.ExpertConsultation{}, id).
		Error; err != nil {
		return err
	}

	return nil
}
