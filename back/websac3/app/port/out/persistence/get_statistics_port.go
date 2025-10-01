package persistence

import (
	"websac3/app/port/out/persistence/db"
)

type GetStatisticsPort interface {
	// Program Lead / Guest statistics
	CountDegreeProgramsByUserID(userID uint, ctx db.Context) (int64, error)
	CountReportsByUserID(userID uint, ctx db.Context) (int64, error)
	CountActiveConsultationsByUserID(userID uint, ctx db.Context) (int64, error)

	// Cybersecurity Auditor statistics
	CountFeedbackReportsByAuditorID(auditorID uint, ctx db.Context) (int64, error)
	CountPendingReportsByAuditorID(auditorID uint, ctx db.Context) (int64, error)
	CountConsultationRequestsByExpertID(expertID uint, ctx db.Context) (int64, error)

	// Admin statistics
	CountActiveUsers(ctx db.Context) (int64, error)
	CountTotalReports(ctx db.Context) (int64, error)
	CountTotalConsultations(ctx db.Context) (int64, error)
	CountTotalAccessRequests(ctx db.Context) (int64, error)
}
