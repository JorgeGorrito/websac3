package seeders

import "websac3/app/port/out/persistence/db"

type professionalRoles struct {
	professionalRoleSeeders []Seeder
}

func ProfessionalRoles() Seeder {
	return &professionalRoles{
		professionalRoleSeeders: []Seeder{
			DevSecOpsRole(),
			CISORole(),
			SecurityArchitectRole(),
			SecurityEngineerRole(),
			SecurityAnalystRole(),
			IncidentResponseSpecialistRole(),
			ThreatHunterRole(),
			MalwareAnalystRole(),
			ForensicInvestigatorRole(),
			PenetrationTesterRole(),
			IAMSpecialistRole(),
			CloudSecuritySpecialistRole(),
			ApplicationSecurityEngineerRole(),
			NetworkSecurityEngineerRole(),
			ComplianceRiskAnalystRole(),
			CybersecurityTrainerAwarenessSpecialistRole(),
		},
	}
}

func (p *professionalRoles) Seed(ctx db.Context) error {
	for _, seeder := range p.professionalRoleSeeders {
		if err := seeder.Seed(ctx); err != nil {
			return err
		}
	}
	return nil
}
