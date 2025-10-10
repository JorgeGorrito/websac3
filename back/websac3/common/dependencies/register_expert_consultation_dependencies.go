package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func RegisterExpertConsultationDependencies(m *manager) {
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateExpertConsultationPort](),
		func() any {
			return repository.NewExpertConsultationRepository()
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateExpertConsultationUseCase](),
		func() any {
			return service.NewCreateExpertConsultationService(
				container.Inject[persistence.CreateExpertConsultationPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[persistence.GetReportPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)

	// Get Expert Consultation Port
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetExpertConsultationPort](),
		func() any {
			return repository.NewExpertConsultationRepository()
		},
	)

	// Get Expert Consultation By ID Use Case
	m.binder.Bind(
		andi.GetAbstractType[usecase.GetExpertConsultationByIDUseCase](),
		func() any {
			return service.NewGetExpertConsultationByIDService(
				container.Inject[persistence.GetExpertConsultationPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	// List User Expert Consultations Use Case
	m.binder.Bind(
		andi.GetAbstractType[usecase.ListUserExpertConsultationsUseCase](),
		func() any {
			return service.NewListUserExpertConsultationsService(
				container.Inject[persistence.GetExpertConsultationPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	// List Pending Expert Consultations Use Case
	m.binder.Bind(
		andi.GetAbstractType[usecase.ListPendingExpertConsultationsUseCase](),
		func() any {
			return service.NewListPendingExpertConsultationsService(
				container.Inject[persistence.GetExpertConsultationPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	// Update Expert Consultation Port
	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateExpertConsultationPort](),
		func() any {
			return repository.NewExpertConsultationRepository()
		},
	)

	// Accept Expert Consultation Use Case
	m.binder.Bind(
		andi.GetAbstractType[usecase.AcceptExpertConsultationUseCase](),
		func() any {
			return service.NewAcceptExpertConsultationService(
				container.Inject[persistence.GetExpertConsultationPort](),
				container.Inject[persistence.UpdateExpertConsultationPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
				container.Inject[template.Provider](),
				container.Inject[notification.SendMailPort](),
			)
		},
	)

	// Reject Expert Consultation Use Case
	m.binder.Bind(
		andi.GetAbstractType[usecase.RejectExpertConsultationUseCase](),
		func() any {
			return service.NewRejectExpertConsultationService(
				container.Inject[persistence.GetExpertConsultationPort](),
				container.Inject[persistence.UpdateExpertConsultationPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
				container.Inject[template.Provider](),
				container.Inject[notification.SendMailPort](),
			)
		},
	)

	// Get Expert Consultation Status Port
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetExpertConsultationStatusPort](),
		func() any {
			return repository.NewExpertConsultationRepository()
		},
	)

	// List Expert Consultation Status Use Case
	m.binder.Bind(
		andi.GetAbstractType[usecase.ListExpertConsultationStatusUseCase](),
		func() any {
			return service.NewListExpertConsultationStatusService(
				container.Inject[persistence.GetExpertConsultationStatusPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)
}
