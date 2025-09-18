package dependencies

import (
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerActivateUserDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[usecase.ActivateUserUseCase](),
		func() any {
			return service.NewActivateUserService(
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.UpdateUserPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)
}
