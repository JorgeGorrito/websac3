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

func (m *manager) registerStatisticsDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetStatisticsPort](),
		func() any {
			return repository.NewStatisticsRepository()
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetUserStatisticsUseCase](),
		func() any {
			return service.NewGetUserStatisticsService(
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetStatisticsPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)
}
