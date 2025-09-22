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

func RegisterCourseTypeDependencies(m *manager) {
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetCourseTypePort](),
		func() any { return repository.NewCourseTypeRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListCourseTypesUseCase](),
		func() any {
			return service.NewListCourseTypesService(
				container.Inject[persistence.GetCourseTypePort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)
}
