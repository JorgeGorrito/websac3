package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerPersonDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.PersonPort](),
		func() any {
			return repository.NewPersonRepository(
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.CreatePersonPort](),
		func() any { return container.Inject[persistence.PersonPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdatePersonPort](),
		func() any { return container.Inject[persistence.PersonPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetPersonPort](),
		func() any { return container.Inject[persistence.PersonPort]() },
	)
}
