package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
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
			)
		},
	)
}
