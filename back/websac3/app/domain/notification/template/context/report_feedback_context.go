package context

type ReportFeedbackContext struct {
	ProgramLeadName   string
	AuditorName       string
	DegreeProgramName string
	ReportScore       float32
	AuditorRating     float32
	GeneralComments   string
	Recommendations   string
	InstitutionName   string
}
