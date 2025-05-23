package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/persistence"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerPersonDependencies() {
	var personRepository *repository.PersonRepository = &repository.PersonRepository{}
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreatePersonPort](),
		func() any { return personRepository },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdatePersonPort](),
		func() any { return personRepository },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetPersonPort](),
		func() any { return personRepository },
	)
}
