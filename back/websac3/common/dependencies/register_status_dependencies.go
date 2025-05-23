package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/persistence"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerStatusDependencies() {
	var statusRepository *repository.StatusRepository = &repository.StatusRepository{}
	m.binder.Bind(
		andi.GetAbstractType[persistence.GetStatusPort](),
		func() any { return statusRepository },
	)
}
