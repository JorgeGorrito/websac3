package dependencies

import (
	"websac3/common/dependencies/container"

	dependencies "github.com/JorgeGorrito/anise-dependency-injection/andi/port/in"
)

type manager struct {
	binder dependencies.Binder
}

var dependenciesManagerInstance *manager = nil

func InitDependenciesManager() {
	container.InitContainer()
	dependenciesManagerInstance = &manager{
		binder: container.GetDependencyBinder(),
	}
	dependenciesManagerInstance.RegisterDependencies()
}

func GetDependenciesManager() *manager {
	if dependenciesManagerInstance == nil {
		InitDependenciesManager()
	}
	return dependenciesManagerInstance
}

func (m *manager) RegisterDependencies() error {
	m.registerPersistenceDependencies()
	m.registerAniseDependencies()
	m.registerAccessRequestDependencies()
	m.registerPersonDependencies()
	m.registerUserDependencies()
	m.registerStatusDependencies()
	return nil
}
