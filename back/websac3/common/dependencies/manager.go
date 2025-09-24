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
	m.registerCommandsFactoryDependencies()
	m.registerMailSenderDependencies()
	m.registerLoggerDependencies()
	m.registerPersistenceDependencies()
	m.registerAniseDependencies()
	m.registerAccessRequestDependencies()
	m.registerExpertConsultationDependencies()
	m.registerPersonDependencies()
	m.registerUserDependencies()
	m.registerStatusDependencies()
	m.registerMediatorDependencies()
	m.registerMessageProviderDependencies()
	m.registerEnumDependencies()
	m.registerTemplateDependencies()
	m.registerPDFDependencies()
	m.registerValidatorDependencies()
	m.registerIdentificationTypeDependencies()
	m.registerHigherEducationInstitutionDependencies()
	m.registerTopicDependencies()
	m.registerDurationUnitDependencies()
	m.registerDegreeProgramDependencies()
	m.registerCourseDependencies()
	m.registerCourseTypeDependencies()
	m.registerCourseNatureDependencies()
	m.registerReportDependencies()
	m.registerProfessionalRoleDependencies()
	m.registerRoleDependencies()
	m.registerApprovedAccessRequestDependencies()
	m.registerRejectedAccessRequestDependencies()
	m.registerDeactivateUserDependencies()
	m.registerActivateUserDependencies()
	return nil
}

func (m *manager) registerCourseTypeDependencies() {
	RegisterCourseTypeDependencies(m)
}

func (m *manager) registerCourseNatureDependencies() {
	RegisterCourseNatureDependencies(m)
}

func (m *manager) registerExpertConsultationDependencies() {
	RegisterExpertConsultationDependencies(m)
}
