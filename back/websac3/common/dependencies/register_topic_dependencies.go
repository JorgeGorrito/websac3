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

func (m *manager) registerTopicDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.TopicPort](),
		func() any { return repository.NewTopicRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetTopicPort](),
		func() any { return container.Inject[persistence.TopicPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListTopicUseCase](),
		func() any {
			return service.NewListTopicService(
				container.Inject[persistence.GetTopicPort](),
				container.Inject[db.Manager](),
				container.Inject[message.Provider](),
			)
		},
	)
}
