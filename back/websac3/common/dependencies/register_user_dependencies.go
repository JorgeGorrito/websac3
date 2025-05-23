package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/persistence"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerUserDependencies() {
	var userRepository *repository.UserRepository = &repository.UserRepository{}
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateUserPort](),
		func() any { return userRepository },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateUserPort](),
		func() any { return userRepository },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetUserPort](),
		func() any { return userRepository },
	)
}
