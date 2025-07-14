package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerAccessRequestDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.AccessRequestPort](),
		func() any {
			return repository.NewAccessRequestRepository(
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateAccessRequestUseCase](),
		func() any {
			return service.NewCreateAccessRequestService(
				container.Inject[persistence.CreateAccessRequestPort](),
				container.Inject[persistence.UpdateAccessRequestPort](),
				container.Inject[persistence.CreatePersonPort](),
				container.Inject[persistence.CreateUserPort](),
				container.Inject[persistence.UpdateUserPort](),
				container.Inject[persistence.UpdatePersonPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetPersonPort](),
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[enum.StatusEnum](),
				container.Inject[message.Provider](),
				container.Inject[notification.SendMailPort](),
				container.Inject[persistence.Manager](),
				container.Inject[template.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateAccessRequestPort](),
		func() any { return container.Inject[persistence.AccessRequestPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateAccessRequestPort](),
		func() any { return container.Inject[persistence.AccessRequestPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetAccessRequestPort](),
		func() any { return container.Inject[persistence.AccessRequestPort]() },
	)
}
