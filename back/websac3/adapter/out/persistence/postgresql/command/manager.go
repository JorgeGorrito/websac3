package command

import "github.com/JorgeGorrito/anise-with-gin/anise/command"

type manager struct{}

func NewCommandManager() *manager {
	return &manager{}
}

func (m *manager) RegisterCommands(registry command.Registry) error {
	registry.Register("migrate:all", NewMigrateAllModels)
	registry.Register("migrate:model", NewMigrateModel)
	registry.Register("seed:run", NewSeedRun)
	return nil
}
