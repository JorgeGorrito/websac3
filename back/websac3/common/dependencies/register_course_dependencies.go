package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	domainusecase "websac3/app/domain/usecase"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerCourseDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateCoursePort](),
		func() any { return repository.NewCourseRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetCoursePort](),
		func() any { return repository.NewCourseRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetCourseTopicsPort](),
		func() any { return repository.NewCourseRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.DeleteCoursePort](),
		func() any { return repository.NewCourseRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateCoursePort](),
		func() any { return repository.NewCourseRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateCourseUseCase](),
		func() any {
			return service.NewCreateCourseService(
				container.Inject[persistence.CreateCoursePort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListCourseByDegreeProgramUseCase](),
		func() any {
			return domainusecase.NewListCourseByDegreeProgramUseCaseImpl(
				service.NewListCourseByDegreeProgramService(
					container.Inject[persistence.GetCoursePort](),
					container.Inject[db.Manager](),
					container.Inject[message.Provider](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListCourseModelsByDegreeProgramUseCase](),
		func() any {
			return domainusecase.NewListCourseModelsByDegreeProgramUseCaseImpl(
				service.NewListCourseModelsByDegreeProgramService(
					container.Inject[persistence.GetCoursePort](),
					container.Inject[db.Manager](),
					container.Inject[message.Provider](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetCourseByIDUseCase](),
		func() any {
			return domainusecase.NewGetCourseByIDUseCaseImpl(
				service.NewGetCourseByIDService(
					container.Inject[persistence.GetCoursePort](),
					container.Inject[db.Manager](),
					container.Inject[message.Provider](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetCourseTopicsByCourseIDUseCase](),
		func() any {
			return domainusecase.NewGetCourseTopicsByCourseIDUseCaseImpl(
				service.NewGetCourseTopicsByCourseIDService(
					container.Inject[persistence.GetCourseTopicsPort](),
					container.Inject[db.Manager](),
					container.Inject[message.Provider](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.DeleteCourseUseCase](),
		func() any {
			return domainusecase.NewDeleteCourseUseCaseImpl(
				service.NewDeleteCourseService(
					container.Inject[persistence.DeleteCoursePort](),
					container.Inject[persistence.GetCoursePort](),
					container.Inject[persistence.GetUserPort](),
					container.Inject[db.Manager](),
					container.Inject[message.Provider](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.UpdateCourseUseCase](),
		func() any {
			return domainusecase.NewUpdateCourseUseCaseImpl(
				service.NewUpdateCourseService(
					container.Inject[persistence.UpdateCoursePort](),
					container.Inject[persistence.GetCoursePort](),
					container.Inject[persistence.GetUserPort](),
					container.Inject[db.Manager](),
					container.Inject[message.Provider](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.BulkCreateCourseUseCase](),
		func() any {
			return domainusecase.NewBulkCreateCourseUseCaseImpl(
				service.NewBulkCreateCourseService(
					container.Inject[db.Manager](),
					container.Inject[persistence.CreateCoursePort](),
					container.Inject[validator.Validator](),
				),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateCourseTopicPort](),
		func() any { return repository.NewCourseRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.BulkCreateCourseTopicUseCase](),
		func() any {
			return domainusecase.NewBulkCreateCourseTopicUseCaseImpl(
				service.NewBulkCreateCourseTopicService(
					container.Inject[db.Manager](),
					container.Inject[persistence.CreateCourseTopicPort](),
					container.Inject[persistence.GetCoursePort](),
					container.Inject[persistence.GetUserPort](),
					container.Inject[validator.Validator](),
				),
			)
		},
	)
}
