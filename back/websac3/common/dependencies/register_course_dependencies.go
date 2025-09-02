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

func (m *manager) registerCourseDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateCoursePort](),
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
}
