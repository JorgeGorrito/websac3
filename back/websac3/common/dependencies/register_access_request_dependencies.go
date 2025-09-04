package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerAccessRequestDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.AccessRequestPort](),
		func() any {
			return repository.NewAccessRequestRepository(
				container.Inject[enum.StatusEnum](),
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
				container.Inject[persistence.UpdatePersonPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetPersonPort](),
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[enum.StatusEnum](),
				container.Inject[message.Provider](),
				container.Inject[notification.SendMailPort](),
				container.Inject[db.Manager](),
				container.Inject[template.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ValidateEmailUseCase](),
		func() any {
			return service.NewValidateEmailService(
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[persistence.UpdateAccessRequestPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
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

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListAccessRequestUseCase](),
		func() any {
			return service.NewListAccessRequestService(
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ApproveAccessRequestUseCase](),
		func() any {
			return service.NewApproveAccessRequestService(
				container.Inject[enum.StatusEnum](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[persistence.UpdateAccessRequestPort](),
				container.Inject[message.Provider](),
				container.Inject[notification.SendMailPort](),
				container.Inject[db.Manager](),
				container.Inject[template.Provider](),
				container.Inject[enum.RoleEnum](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.RejectAccessRequestUseCase](),
		func() any {
			return service.NewRejectAccessRequestService(
				container.Inject[enum.StatusEnum](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[persistence.UpdateAccessRequestPort](),
				container.Inject[message.Provider](),
				container.Inject[notification.SendMailPort](),
				container.Inject[db.Manager](),
				container.Inject[template.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateUserFromTokenUseCase](),
		func() any {
			return service.NewCreateUserFromTokenService(
				container.Inject[persistence.GetAccessRequestPort](),
				container.Inject[persistence.CreateUserPort](),
				container.Inject[persistence.CreatePersonPort](),
				container.Inject[persistence.UpdateAccessRequestPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[enum.RoleEnum](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)
}
