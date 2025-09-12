package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/pdf"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"

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
		andi.GetAbstractType[usecase.CreateDegreeProgramUseCase](),
		func() any {
			return service.NewCreateDegreeProgramService(
				container.Inject[persistence.CreateDegreeProgramPort](),
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
}
