package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerUserDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.UserPort](),
		func() any {
			return repository.NewUserRepository(
				container.Inject[message.Provider](),
			)
		},
	)
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateUserPort](),
		func() any { return container.Inject[persistence.UserPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateUserPort](),
		func() any { return container.Inject[persistence.UserPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetUserPort](),
		func() any { return container.Inject[persistence.UserPort]() },
	)
}
