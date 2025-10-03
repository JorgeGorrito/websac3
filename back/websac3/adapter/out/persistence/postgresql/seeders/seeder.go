package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	_db "websac3/app/port/out/persistence/db"
)

type Seeder interface {
	Seed(_db.Context) error
}

type NewSeeder func() Seeder

var registry map[string]NewSeeder = map[string]NewSeeder{
	"essential_data":                                  EssentialData,
	"devsecops_role":                                  DevSecOpsRole,
	"ciso_role":                                       CISORole,
	"security_architect_role":                         SecurityArchitectRole,
	"security_engineer_role":                          SecurityEngineerRole,
	"security_analyst_role":                           SecurityAnalystRole,
	"incident_response_specialist_role":               IncidentResponseSpecialistRole,
	"threat_hunter_role":                              ThreatHunterRole,
	"malware_analyst_role":                            MalwareAnalystRole,
	"forensic_investigator_role":                      ForensicInvestigatorRole,
	"penetration_tester_role":                         PenetrationTesterRole,
	"iam_specialist_role":                             IAMSpecialistRole,
	"cloud_security_specialist_role":                  CloudSecuritySpecialistRole,
	"application_security_engineer_role":              ApplicationSecurityEngineerRole,
	"network_security_engineer_role":                  NetworkSecurityEngineerRole,
	"compliance_risk_analyst_role":                    ComplianceRiskAnalystRole,
	"cybersecurity_trainer_awareness_specialist_role": CybersecurityTrainerAwarenessSpecialistRole,
	"professional_roles":                              ProfessionalRoles,
}

func GetSeederConstructorByName(name string) NewSeeder {
	return registry[name]
}

func ResetAutoIncrement(ctx _db.Context, tableName string) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}
	sql := fmt.Sprintf(`
		SELECT setval(
			pg_get_serial_sequence('%s', 'id'),
			COALESCE(MAX(id), 0) + 1,
			false
		) FROM %s;
	`, tableName, tableName)

	return dbCtx.DB().Exec(sql).Error
}
