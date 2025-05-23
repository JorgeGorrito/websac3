package container

import (
	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	dependencies "github.com/JorgeGorrito/anise-dependency-injection/andi/port/in"
)

var containerInstance *andi.Container = nil

func InitContainer() {
	if containerInstance == nil {
		containerInstance = andi.NewContainer()
	}
}

func GetDependencyBinder() dependencies.Binder {
	return containerInstance
}

func Inject[T any]() T {
	return andi.Inject[T](containerInstance)
}
