package dependencies

import "github.com/JorgeGorrito/anise-with-gin/anise/dependencies"

type manager struct{}

func (m *manager) RegisterDependencies(dependenciesBinder dependencies.Binder) error {
	return nil
}

var instance *manager = nil

func GetManager() *manager {
	if instance == nil {
		instance = &manager{}
	}
	return instance
}
