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

func RegisterExpertConsultationDependencies(m *manager) {
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateExpertConsultationPort](),
		func() any {
			return repository.NewExpertConsultationRepository()
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.CreateExpertConsultationUseCase](),
		func() any {
			return service.NewCreateExpertConsultationService(
				container.Inject[persistence.CreateExpertConsultationPort](),
				container.Inject[persistence.GetUserPort](),
				container.Inject[persistence.GetDegreeProgramPort](),
				container.Inject[persistence.GetReportPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)
}
