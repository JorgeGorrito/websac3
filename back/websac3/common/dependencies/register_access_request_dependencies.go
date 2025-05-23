package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerAccessRequestDependencies() {
	var accessRequestRepository *repository.AccessRequestRepository = repository.NewAccessRequestRepository()
	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateAccessRequestUseCase](),
		func() any {
			return service.NewCreateAccessRequestService(
				container.Inject[persistence.CreateAccessRequestPort](),
				container.Inject[persistence.CreatePersonPort](),
				container.Inject[persistence.CreateUserPort](),
				container.Inject[persistence.UpdateUserPort](),
				container.Inject[persistence.UpdatePersonPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetPersonPort](),
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[persistence.GetStatusPort](),
				container.Inject[persistence.TransactionManager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateAccessRequestPort](),
		func() any { return accessRequestRepository },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetAccessRequestPort](),
		func() any { return accessRequestRepository },
	)
}
