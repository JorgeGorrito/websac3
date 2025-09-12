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
					"access_request_confirmation:es": atemplate.NewAccessRequestConfirmationTemplate(
						baseDir + "es/access_request_confirmation.html",
					),
					"access_request_approved:es": atemplate.NewAccessRequestApprovedTemplate(
						baseDir + "es/access_request_approved.html",
					),
					"access_request_rejected:es": atemplate.NewAccessRequestRejectedTemplate(
						baseDir + "es/access_request_rejected.html",
					),
					"access_request_confirmation:en": atemplate.NewAccessRequestConfirmationTemplate(
						baseDir + "en/access_request_confirmation.html",
					),
					"access_request_approved:en": atemplate.NewAccessRequestApprovedTemplate(
						baseDir + "en/access_request_approved.html",
					),
					"access_request_rejected:en": atemplate.NewAccessRequestRejectedTemplate(
						baseDir + "en/access_request_rejected.html",
					),
					"report:es": atemplate.NewReportTemplate(
						baseDir + "es/report.html",
					),
					"report:en": atemplate.NewReportTemplate(
						baseDir + "en/report.html",
					),
					"degree_program_evaluation:es": atemplate.NewDegreeProgramEvaluationTemplate(
						baseDir + "es/degree_program_evaluation.html",
					),
					"degree_program_evaluation:en": atemplate.NewDegreeProgramEvaluationTemplate(
						baseDir + "en/degree_program_evaluation.html",
					),
				},
			)
		},
	)
}
