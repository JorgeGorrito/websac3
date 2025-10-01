package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type StatisticsRepository struct {
	Repository
}

func NewStatisticsRepository() *StatisticsRepository {
	return &StatisticsRepository{}
}

// CountDegreeProgramsByUserID cuenta los programas de grado creados por un usuario
func (r *StatisticsRepository) CountDegreeProgramsByUserID(userID uint, ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Where("created_by = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountReportsByUserID cuenta los reportes generados para programas del usuario
func (r *StatisticsRepository) CountReportsByUserID(userID uint, ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Joins("JOIN degree_programs ON reports.degree_program_id = degree_programs.id").
		Where("degree_programs.created_by = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountActiveConsultationsByUserID cuenta las asesorías activas solicitadas por el usuario
func (r *StatisticsRepository) CountActiveConsultationsByUserID(userID uint, ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	// Asesorías activas: las que están en estado "accepted"
	if err := dbCtx.DB().
		Model(&model.ExpertConsultation{}).
		Joins("JOIN expert_consultation_statuses ON expert_consultations.status_id = expert_consultation_statuses.id").
		Joins("JOIN expert_consultation_status_names ON expert_consultation_statuses.id = expert_consultation_status_names.expert_consultation_status_id").
		Where("expert_consultations.requester_id = ? AND expert_consultation_status_names.name = ?", userID, "accepted").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountFeedbackReportsByAuditorID cuenta los reportes retroalimentados por un auditor
func (r *StatisticsRepository) CountFeedbackReportsByAuditorID(auditorID uint, ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.ReportFeedback{}).
		Where("auditor_id = ?", auditorID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountPendingReportsByAuditorID cuenta los reportes pendientes de retroalimentación
// Para un auditor, contamos todos los reportes que aún no tienen retroalimentación
func (r *StatisticsRepository) CountPendingReportsByAuditorID(auditorID uint, ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	// Reportes que no tienen retroalimentación
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Where("id NOT IN (SELECT report_id FROM report_feedbacks)").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountConsultationRequestsByExpertID cuenta las solicitudes de asesoría asignadas a un experto
func (r *StatisticsRepository) CountConsultationRequestsByExpertID(expertID uint, ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.ExpertConsultation{}).
		Where("expert_id = ?", expertID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountActiveUsers cuenta los usuarios activos en el sistema
func (r *StatisticsRepository) CountActiveUsers(ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.User{}).
		Where("deactivated_at IS NULL").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountTotalReports cuenta el total de reportes en el sistema
func (r *StatisticsRepository) CountTotalReports(ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountTotalConsultations cuenta el total de consultas de expertos en el sistema
func (r *StatisticsRepository) CountTotalConsultations(ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.ExpertConsultation{}).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// CountTotalAccessRequests cuenta el total de solicitudes de acceso en el sistema
func (r *StatisticsRepository) CountTotalAccessRequests(ctx _db.Context) (int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := dbCtx.DB().
		Model(&model.AccessRequest{}).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
