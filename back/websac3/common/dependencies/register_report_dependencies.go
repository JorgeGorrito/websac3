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

func (m *manager) registerReportDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateReportPort](),
		func() any { return repository.NewReportRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetReportPort](),
		func() any { return repository.NewReportRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListReportsByDegreeProgramUseCase](),
		func() any {
			return service.NewListReportsByDegreeProgramService(
				container.Inject[persistence.GetReportPort](),
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetReportByIDUseCase](),
		func() any {
			return service.NewGetReportByIDService(
				container.Inject[persistence.GetReportPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)

	// Report Feedback Dependencies
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateReportFeedbackPort](),
		func() any { return repository.NewReportFeedbackRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetReportFeedbackPort](),
		func() any { return repository.NewReportFeedbackRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.ListReportsPendingFeedbackPort](),
		func() any { return repository.NewReportFeedbackRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.ListReportFeedbacksPort](),
		func() any { return repository.NewReportFeedbackRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateReportFeedbackUseCase](),
		func() any {
			return service.NewCreateReportFeedbackService(
				container.Inject[persistence.GetReportPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.CreateReportFeedbackPort](),
				container.Inject[persistence.GetReportFeedbackPort](),
				container.Inject[notification.SendMailPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
				container.Inject[template.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetReportFeedbackUseCase](),
		func() any {
			return service.NewGetReportFeedbackService(
				container.Inject[persistence.GetReportFeedbackPort](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListReportsPendingFeedbackUseCase](),
		func() any {
			return service.NewListReportsPendingFeedbackService(
				container.Inject[persistence.ListReportsPendingFeedbackPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListReportFeedbacksUseCase](),
		func() any {
			return service.NewListReportFeedbacksService(
				container.Inject[persistence.ListReportFeedbacksPort](),
				container.Inject[db.Manager](),
			)
		},
	)
}
