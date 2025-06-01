package dependencies

import (
	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	"websac3/common/mediator"
)

func (m *manager) registerMediatorDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[mediator.IMediator](),
		func() any {
			return mediator.New()
		},
	)
}
