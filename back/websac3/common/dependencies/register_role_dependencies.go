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

func (m *manager) registerRoleDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[usecase.ListRoleUseCase](),
		func() any {
			return service.NewListRoleService(
				container.Inject[persistence.ListRolePort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.ListRolePort](),
		func() any {
			return repository.NewRoleRepository()
		},
	)
}
