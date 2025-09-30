package repository

import (
	"errors"
	"time"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	commonfilter "websac3/common/filter"
	"websac3/common/mapper"
	"websac3/common/paginator"
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
		Preload("Requester.Person.HigherEducationInstitution").
		Preload("Requester.Person.HigherEducationInstitution.Ownership").
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

// GetByUserID obtiene las solicitudes de asesoría de un usuario específico
func (e *ExpertConsultationRepository) GetByUserID(userID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, ctx _db.Context) ([]entity.ExpertConsultation, uint, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Query base para obtener solicitudes de asesoría del usuario
	query := dbCtx.DB().Model(&model.ExpertConsultation{}).
		Where("requester_id = ?", userID)

	// Aplicar filtros básicos usando subqueries para evitar problemas con JOINs
	for field, filterData := range filters {
		for operator, value := range filterData {
			if value == "" {
				continue
			}

			if field == "status_id" {
				if operator == "eq" {
					query = query.Where("status_id = ?", value)
				}
			} else if field == "Status.name" {
				if operator == "eq" {
					query = query.Where("status_id IN (SELECT ecs.id FROM expert_consultation_statuses ecs JOIN expert_consultation_status_names ecsn ON ecs.id = ecsn.expert_consultation_status_id WHERE ecsn.name = ?)", value)
				} else if operator == "cont" {
					query = query.Where("status_id IN (SELECT ecs.id FROM expert_consultation_statuses ecs JOIN expert_consultation_status_names ecsn ON ecs.id = ecsn.expert_consultation_status_id WHERE ecsn.name ILIKE ?)", "%"+value+"%")
				}
			} else if field == "DegreeProgram.name" {
				if operator == "eq" {
					query = query.Where("degree_program_id IN (SELECT id FROM degree_programs WHERE name = ?)", value)
				} else if operator == "cont" {
					query = query.Where("degree_program_id IN (SELECT id FROM degree_programs WHERE name ILIKE ?)", "%"+value+"%")
				}
			}
		}
	}

	// Contar total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación y ordenamiento con Preloads
	var expertConsultations []model.ExpertConsultation
	if err := query.
		Preload("Requester").
		Preload("Requester.Person").
		Preload("DegreeProgram").
		Preload("Report").
		Preload("Expert").
		Preload("Expert.Person").
		Preload("Status").
		Preload("Status.Names").
		Order("created_at DESC").
		Offset(int((paginationParams.Currentpage - 1) * paginationParams.ItemsPerpage)).
		Limit(int(paginationParams.ItemsPerpage)).
		Find(&expertConsultations).
		Error; err != nil {
		return nil, 0, err
	}

	// Mapear a entidades
	var results []entity.ExpertConsultation
	for _, expertConsultation := range expertConsultations {
		var expertConsultationEntity entity.ExpertConsultation
		if expertConsultationEntity, err = mapper.Map[model.ExpertConsultation, entity.ExpertConsultation](&expertConsultation); err != nil {
			return nil, 0, err
		}
		results = append(results, expertConsultationEntity)
	}

	return results, uint(total), nil
}

// GetPendingConsultations obtiene las solicitudes de asesoría pendientes
func (e *ExpertConsultationRepository) GetPendingConsultations(paginationParams paginator.PaginationParams, filters commonfilter.Params, ctx _db.Context) ([]entity.ExpertConsultation, uint, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Query base para obtener solicitudes de asesoría pendientes (status_id = 1)
	query := dbCtx.DB().Model(&model.ExpertConsultation{}).
		Preload("Requester").
		Preload("Requester.Person").
		Preload("Requester.Person.HigherEducationInstitution").
		Preload("Requester.Person.HigherEducationInstitution.Ownership").
		Preload("DegreeProgram").
		Preload("Report").
		Preload("Status").
		Preload("Status.Names").
		Where("status_id = ?", 1) // 1 = pending status

	// Aplicar filtros básicos
	for field, filterData := range filters {
		for operator, value := range filterData {
			if value == "" {
				continue
			}

			if field == "Requester.Person.name" {
				if operator == "eq" {
					query = query.Joins("JOIN users ON expert_consultations.requester_id = users.id").
						Joins("JOIN people ON users.person_id = people.id").
						Where("people.name = ?", value)
				} else if operator == "cont" {
					query = query.Joins("JOIN users ON expert_consultations.requester_id = users.id").
						Joins("JOIN people ON users.person_id = people.id").
						Where("people.name ILIKE ?", "%"+value+"%")
				}
			} else if field == "Requester.Person.lastname" {
				if operator == "eq" {
					query = query.Joins("JOIN users ON expert_consultations.requester_id = users.id").
						Joins("JOIN people ON users.person_id = people.id").
						Where("people.lastname = ?", value)
				} else if operator == "cont" {
					query = query.Joins("JOIN users ON expert_consultations.requester_id = users.id").
						Joins("JOIN people ON users.person_id = people.id").
						Where("people.lastname ILIKE ?", "%"+value+"%")
				}
			} else if field == "DegreeProgram.name" {
				if operator == "eq" {
					query = query.Joins("JOIN degree_programs ON expert_consultations.degree_program_id = degree_programs.id").
						Where("degree_programs.name = ?", value)
				} else if operator == "cont" {
					query = query.Joins("JOIN degree_programs ON expert_consultations.degree_program_id = degree_programs.id").
						Where("degree_programs.name ILIKE ?", "%"+value+"%")
				}
			}
		}
	}

	// Contar total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación y ordenamiento
	var expertConsultations []model.ExpertConsultation
	if err := query.
		Order("created_at DESC").
		Offset(int((paginationParams.Currentpage - 1) * paginationParams.ItemsPerpage)).
		Limit(int(paginationParams.ItemsPerpage)).
		Find(&expertConsultations).
		Error; err != nil {
		return nil, 0, err
	}

	// Mapear a entidades
	var results []entity.ExpertConsultation
	for _, expertConsultation := range expertConsultations {
		var expertConsultationEntity entity.ExpertConsultation
		if expertConsultationEntity, err = mapper.Map[model.ExpertConsultation, entity.ExpertConsultation](&expertConsultation); err != nil {
			return nil, 0, err
		}
		results = append(results, expertConsultationEntity)
	}

	return results, uint(total), nil
}

// AcceptConsultation acepta una solicitud de asesoría
func (e *ExpertConsultationRepository) AcceptConsultation(consultationID uint, expertResponse string, expertID uint, ctx _db.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	result := dbCtx.DB().Model(&model.ExpertConsultation{}).
		Where("id = ? AND status_id = ? AND expert_id IS NULL", consultationID, 1).
		Updates(map[string]interface{}{
			"expert_response": expertResponse,
			"expert_id":       expertID,
			"status_id":       3, // 3 = accepted
			"answered_at":     &now,
			"updated_at":      now,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se pudo aceptar la consulta - puede que ya esté asignada o no esté pendiente")
	}

	return nil
}

// RejectConsultation rechaza una solicitud de asesoría
func (e *ExpertConsultationRepository) RejectConsultation(consultationID uint, expertResponse string, expertID uint, ctx _db.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	result := dbCtx.DB().Model(&model.ExpertConsultation{}).
		Where("id = ? AND status_id = ? AND expert_id IS NULL", consultationID, 1).
		Updates(map[string]interface{}{
			"expert_response": expertResponse,
			"expert_id":       expertID,
			"status_id":       2, // 2 = rejected
			"answered_at":     &now,
			"updated_at":      now,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se pudo rechazar la consulta - puede que ya esté asignada o no esté pendiente")
	}

	return nil
}

// GetAll obtiene todos los estados de asesorías de experto con paginación y filtros
func (e *ExpertConsultationRepository) GetAll(paginationParams paginator.PaginationParams, filters commonfilter.Params, ctx _db.Context) ([]entity.ExpertConsultationStatus, uint, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Query base para obtener estados de asesorías de experto
	query := dbCtx.DB().Model(&model.ExpertConsultationStatus{})

	// Aplicar filtros básicos usando subqueries para evitar problemas con JOINs
	for field, filterData := range filters {
		for operator, value := range filterData {
			if value == "" {
				continue
			}

			if field == "Names.name" {
				if operator == "eq" {
					query = query.Where("id IN (SELECT expert_consultation_status_id FROM expert_consultation_status_names WHERE name = ?)", value)
				} else if operator == "cont" {
					query = query.Where("id IN (SELECT expert_consultation_status_id FROM expert_consultation_status_names WHERE name ILIKE ?)", "%"+value+"%")
				}
			} else if field == "Names.lang" {
				if operator == "eq" {
					query = query.Where("id IN (SELECT expert_consultation_status_id FROM expert_consultation_status_names WHERE lang = ?)", value)
				}
			}
		}
	}

	// Contar total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación y ordenamiento con Preloads
	var expertConsultationStatuses []model.ExpertConsultationStatus
	if err := query.
		Preload("Names").
		Order("id ASC").
		Offset(int((paginationParams.Currentpage - 1) * paginationParams.ItemsPerpage)).
		Limit(int(paginationParams.ItemsPerpage)).
		Find(&expertConsultationStatuses).
		Error; err != nil {
		return nil, 0, err
	}

	// Mapear a entidades
	var results []entity.ExpertConsultationStatus
	for _, status := range expertConsultationStatuses {
		var statusEntity entity.ExpertConsultationStatus
		if statusEntity, err = mapper.Map[model.ExpertConsultationStatus, entity.ExpertConsultationStatus](&status); err != nil {
			return nil, 0, err
		}
		results = append(results, statusEntity)
	}

	return results, uint(total), nil
}
