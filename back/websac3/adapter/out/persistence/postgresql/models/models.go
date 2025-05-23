package models

type NewBaseModel func() any

var registry map[string]NewBaseModel = map[string]NewBaseModel{
	"users":                         func() any { return &User{} },
	"roles":                         func() any { return &Role{} },
	"permissions":                   func() any { return &Permission{} },
	"access_requests":               func() any { return &AccessRequest{} },
	"access_request_statuses":       func() any { return &Status{} },
	"people":                        func() any { return &Person{} },
	"municipalities":                func() any { return &Municipality{} },
	"departments":                   func() any { return &Department{} },
	"institutional_categories":      func() any { return &InstitutionalCategory{} },
	"ownerships":                    func() any { return &Ownership{} },
	"higher_education_institutions": func() any { return &HigherEducationInstitution{} },
	"identification_types":          func() any { return &IdentificationType{} },
}

func GetRegistryAllConstructModelBase() map[string]NewBaseModel {
	return registry
}

func GetConstructModelBaseByName(name string) NewBaseModel {
	return registry[name]
}
