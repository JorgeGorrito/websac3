package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	domainusecase "websac3/app/domain/usecase"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/pdf"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerDegreeProgramDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetProfessionalRolePort](),
		func() any { return repository.NewProfessionalRoleRepository() },
	)
	m.binder.Bind(
		andi.GetAbstractType[usecase.EvaluateDegreeProgramUseCase](),
		func() any {
			return service.NewEvaluateDegreeProgramService(
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[persistence.GetProfessionalRolePort](),
				container.Inject[persistence.CreateReportPort](),
				container.Inject[usecase.GetUserByIDUseCase](),
				container.Inject[message.Provider](),
				container.Inject[template.Provider](),
				container.Inject[pdf.Converter](),
				container.Inject[notification.SendMailPort](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.DegreeProgramPort](),
		func() any { return repository.NewDegreeProgramRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateDegreeProgramPort](),
		func() any { return container.Inject[persistence.DegreeProgramPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetDegreeProgramPort](),
		func() any { return container.Inject[persistence.DegreeProgramPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateDegreeProgramPort](),
		func() any { return container.Inject[persistence.DegreeProgramPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.DeleteDegreeProgramPort](),
		func() any { return container.Inject[persistence.DegreeProgramPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateDegreeProgramUseCase](),
		func() any {
			return service.NewCreateDegreeProgramService(
				container.Inject[persistence.CreateDegreeProgramPort](),
				container.Inject[persistence.GetProfessionalRolePort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListDegreeProgramUseCase](),
		func() any {
			return service.NewListDegreeProgramService(
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[enum.RoleEnum](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetDegreeProgramByIDUseCase](),
		func() any {
			return service.NewGetDegreeProgramByIDService(
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.UpdateDegreeProgramUseCase](),
		func() any {
			return service.NewUpdateDegreeProgramService(
				container.Inject[persistence.UpdateDegreeProgramPort](),
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[persistence.GetProfessionalRolePort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.BulkCreateDegreeProgramUseCase](),
		func() any {
			return domainusecase.NewBulkCreateDegreeProgramUseCaseImpl(
				service.NewBulkCreateDegreeProgramService(
					container.Inject[db.Manager](),
					container.Inject[persistence.CreateDegreeProgramPort](),
					container.Inject[persistence.GetProfessionalRolePort](),
					container.Inject[persistence.GetDurationUnitPort](),
					container.Inject[persistence.GetFormationLevelPort](),
					container.Inject[validator.Validator](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.DeleteDegreeProgramUseCase](),
		func() any {
			return service.NewDeleteDegreeProgramService(
				container.Inject[db.Manager](),
				container.Inject[persistence.DeleteDegreeProgramPort](),
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[message.Provider](),
			)
		},
	)
}
