package model

type NewBaseModel func() any

var registry map[string]NewBaseModel = map[string]NewBaseModel{
	"users":                             func() any { return &User{} },
	"roles":                             func() any { return &Role{} },
	"access_requests":                   func() any { return &AccessRequest{} },
	"access_request_statuses":           func() any { return &Status{} },
	"people":                            func() any { return &Person{} },
	"municipalities":                    func() any { return &Municipality{} },
	"departments":                       func() any { return &Department{} },
	"institutional_categories":          func() any { return &InstitutionalCategory{} },
	"ownerships":                        func() any { return &Ownership{} },
	"higher_education_institutions":     func() any { return &HigherEducationInstitution{} },
	"identification_types":              func() any { return &IdentificationType{} },
	"emails":                            func() any { return &Email{} },
	"modules":                           func() any { return &Module{} },
	"actions":                           func() any { return &Action{} },
	"permissions":                       func() any { return &Permission{} },
	"knowledge_area_names":              func() any { return &KnowledgeAreaName{} },
	"knowledge_areas":                   func() any { return &KnowledgeArea{} },
	"topic_names":                       func() any { return &TopicName{} },
	"topics":                            func() any { return &Topic{} },
	"duration_units":                    func() any { return &DurationUnit{} },
	"duration_unit_names":               func() any { return &DurationUnitName{} },
	"degree_programs":                   func() any { return &DegreeProgram{} },
	"course_types":                      func() any { return &CourseType{} },
	"course_type_names":                 func() any { return &CourseTypeName{} },
	"course_natures":                    func() any { return &CourseNature{} },
	"course_nature_names":               func() any { return &CourseNatureName{} },
	"courses":                           func() any { return &Course{} },
	"course_topics":                     func() any { return &CourseTopic{} },
	"professional_roles":                func() any { return &ProfessionalRole{} },
	"professional_role_knowledge_areas": func() any { return &ProfessionalRoleKnowledgeArea{} },
	"professional_role_knowledge_area_topics": func() any { return &ProfessionalRoleKnowledgeAreaTopic{} },
	"reports":                           func() any { return &Report{} },
	"knowledge_area_reports":            func() any { return &KnowledgeAreaReport{} },
	"topic_reports":                     func() any { return &TopicReport{} },
	"unexpected_knowledge_area_reports": func() any { return &UnexpectedKnowledgeAreaReport{} },
	"unexpected_topic_reports":          func() any { return &UnexpectedTopicReport{} },
	"report_feedbacks":                  func() any { return &ReportFeedback{} },
	"knowledge_area_feedbacks":          func() any { return &KnowledgeAreaFeedback{} },
}

func GetRegistryAllConstructModelBase() map[string]NewBaseModel {
	return registry
}

func GetConstructModelBaseByName(name string) NewBaseModel {
	return registry[name]
}
