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

func (m *manager) registerIdentificationTypeDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[usecase.ListIdentificationTypeUseCase](),
		func() any {
			return service.NewListIdentificationTypeService(
				container.Inject[persistence.GetIdentificationTypePort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.IdentificationTypePort](),
		func() any {
			return repository.NewIdentificationTypeRepository()
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetIdentificationTypePort](),
		func() any { return container.Inject[persistence.IdentificationTypePort]() },
	)
}
