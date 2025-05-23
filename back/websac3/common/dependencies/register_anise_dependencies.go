package dependencies

import (
	"websac3/adapter/in/web/config"
	"websac3/adapter/in/web/routing"
	"websac3/adapter/out/persistence/postgresql/command"

	acmd "github.com/JorgeGorrito/anise-with-gin/anise/command"
	aconf "github.com/JorgeGorrito/anise-with-gin/anise/config"
	arou "github.com/JorgeGorrito/anise-with-gin/anise/routing"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerAniseDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[aconf.Manager](),
		func() any { return config.NewConfigManager() },
	)

	m.binder.Bind(
		andi.GetAbstractType[arou.Manager](),
		func() any { return routing.NewRoutingManager() },
	)

	m.binder.Bind(
		andi.GetAbstractType[acmd.Manager](),
		func() any { return command.NewCommandManager() },
	)
}
