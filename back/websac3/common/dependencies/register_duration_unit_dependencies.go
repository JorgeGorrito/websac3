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

func (m *manager) registerDurationUnitDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.DurationUnitPort](),
		func() any { return repository.NewDurationUnitRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetDurationUnitPort](),
		func() any { return container.Inject[persistence.DurationUnitPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListDurationUnitUseCase](),
		func() any {
			return service.NewListDurationUnitService(
				container.Inject[persistence.GetDurationUnitPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)
}
