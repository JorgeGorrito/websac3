package dependencies

import (
	"websac3/app/port/out/message"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerValidatorDependencies() error {
	m.binder.Bind(
		andi.GetAbstractType[validator.Validator](),
		func() any {
			return validator.New(
				container.Inject[message.Provider](),
			)
		},
	)
	return nil
}
