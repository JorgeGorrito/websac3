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
}
