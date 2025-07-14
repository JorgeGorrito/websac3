package dependencies

import (
	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

func (m *manager) registerCommandsFactoryDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[command.Factory](),
		func() any { return command.NewFactory() },
	)
}
