package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/persistence"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerReportDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateReportPort](),
		func() any { return repository.NewReportRepository() },
	)
}
