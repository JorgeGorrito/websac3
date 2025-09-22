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

func RegisterCourseNatureDependencies(m *manager) {
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetCourseNaturePort](),
		func() any { return repository.NewCourseNatureRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListCourseNaturesUseCase](),
		func() any {
			return service.NewListCourseNaturesService(
				container.Inject[persistence.GetCourseNaturePort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)
}
