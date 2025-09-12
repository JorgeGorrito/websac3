package dependencies

import (
	"websac3/adapter/out/pdf/wkhtml"
	appdf "websac3/app/port/out/pdf"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerPDFDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[appdf.Converter](),
		func() any { return wkhtml.New() },
	)
}
