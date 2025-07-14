package dependencies

import (
	"os"
	"websac3/app/port/out/notification/template"

	atemplate "websac3/adapter/out/notification/template"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerTemplateDependencies() {
	var baseDir = os.Getenv("TEMPLATES_BASE_DIR")
	m.binder.Bind(
		andi.GetAbstractType[template.Provider](),
		func() any {
			return atemplate.NewProvider(
				map[string]template.Template{
					"access_request_confirmation": atemplate.NewAccessRequestConfirmationTemplate(
						baseDir + "access_request_confirmation.html",
					),
				},
			)
		},
	)
}
