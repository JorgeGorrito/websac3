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

func (m *manager) registerFormationLevelDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.FormationLevelPort](),
		func() any { return repository.NewFormationLevelRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetFormationLevelPort](),
		func() any { return container.Inject[persistence.FormationLevelPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListFormationLevelUseCase](),
		func() any {
			return service.NewListFormationLevelService(
				container.Inject[persistence.GetFormationLevelPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

}
