package command

import (
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

type manager struct {
	factory command.Factory
}

func (m *manager) RegisterCommands(registry command.Registry) error {
	registry.Register("migrate:all", NewMigrateAllModels)
	registry.Register("migrate:model", NewMigrateModel)
	registry.Register("migrate:reset", NewDBReset)

	registry.Register("seed:run", NewSeedRun)
	return nil
}

func NewCommandManager() *manager {
	return &manager{factory: container.Inject[command.Factory]()}
}

func (m *manager) GetRetrieverCommand() command.Retriever {
	return m.factory
}

func (m *manager) GetRegistryCommand() command.Registry {
	return m.factory
}
