package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerRejectedAccessRequestDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[usecase.ListRejectedAccessRequestUseCase](),
		func() any {
			return service.NewListRejectedAccessRequestService(
				container.Inject[persistence.ListRejectedAccessRequestPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.ListRejectedAccessRequestPort](),
		func() any {
			return repository.NewAccessRequestRepository(container.Inject[enum.StatusEnum]())
		},
	)
}
