package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type securityAnalystRole struct{}

func SecurityAnalystRole() Seeder { return &securityAnalystRole{} }

func (s *securityAnalystRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Security Analyst
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Security Analyst").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Security Analyst"}
		if err := dbCtx.DB().Create(&role).Error; err != nil {
			return fmt.Errorf("cannot create professional role: %w", err)
		}
	}

	// Helper: find KA by name (EN) -> ka_id
	findKA := func(name string) (uint, error) {
		var kan model.KnowledgeAreaName
		if err := dbCtx.DB().Where("name = ? AND lang = ?", name, "en").First(&kan).Error; err != nil {
			return 0, fmt.Errorf("knowledge area '%s' not found: %w", name, err)
		}
		return kan.KnowledgeAreaID, nil
	}

	// Helper: find Topic by name (EN) -> topic_id
	findTopic := func(name string) (uint, error) {
		var tn model.TopicName
		if err := dbCtx.DB().Where("name = ? AND lang = ?", name, "en").First(&tn).Error; err != nil {
			return 0, fmt.Errorf("topic '%s' not found: %w", name, err)
		}
		return tn.TopicID, nil
	}

	// 2) Define KA expected for Security Analyst (weights sum to 1.0)
	type topicCfg struct {
		Name  string
		Hours float32
	}
	type kaCfg struct {
		Name   string
		Weight float32
		Topics []topicCfg
	}

	cfg := []kaCfg{
		{
			Name: "Organizational Security", Weight: 0.35,
			Topics: []topicCfg{
				{Name: "Security Program Management - Security Metrics", Hours: 8},
				{Name: "Security Operations - Security Convergence", Hours: 10},
				{Name: "Security Operations - Global Security Operations Centers (GSOCs)", Hours: 10},
				{Name: "Security Program Management - Project Management", Hours: 6},
				{Name: "Security Program Management - Resource Management", Hours: 8},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 10},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Disaster Recovery", Hours: 6},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Business Continuity", Hours: 6},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 6},
				{Name: "Security Program Management - Project Management", Hours: 4},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 6},
				{Name: "Risk Management - Risk Identification", Hours: 6},
				{Name: "Risk Management - Risk Assessment and Analysis", Hours: 8},
				{Name: "Risk Management - Insider Threats", Hours: 6},
				{Name: "Risk Management - Risk Measurement and Evaluation Models and Methodologies", Hours: 6},
				{Name: "Risk Management - Risk Control", Hours: 6},
			},
		},
		{
			Name: "System Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "System Control - Intrusion Detection", Hours: 10},
				{Name: "System Control - Attacks", Hours: 8},
				{Name: "System Control - Defenses", Hours: 6},
				{Name: "System Control - Auditing", Hours: 8},
				{Name: "System Control - Malware", Hours: 8},
				{Name: "System Control - Vulnerability Models", Hours: 8},
				{Name: "System Control - Penetration Testing", Hours: 6},
				{Name: "System Control - Forensics", Hours: 8},
				{Name: "System Control - Recovery and Resilience", Hours: 4},
				{Name: "Systems Management - Policy Models", Hours: 6},
				{Name: "Systems Management - Policy Composition", Hours: 4},
				{Name: "Systems Management - Use of Automation", Hours: 4},
				{Name: "Systems Management - Patching and Vulnerability Life Cycle", Hours: 6},
				{Name: "Systems Management - Operation", Hours: 6},
				{Name: "Systems Management - Insider Threat", Hours: 6},
				{Name: "System Testing - Requirements Validation", Hours: 4},
				{Name: "System Testing - Component Composition Validation", Hours: 4},
				{Name: "System Testing - Unit Testing versus System Testing", Hours: 4},
			},
		},
		{
			Name: "Connection Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Network Defense - Network Monitoring", Hours: 10},
				{Name: "Network Defense - Network Traffic Analysis", Hours: 10},
				{Name: "Network Defense - Threat Hunting and Machine Learning", Hours: 8},
				{Name: "Network Defense - IDS/IPS Implementation", Hours: 8},
				{Name: "Network Defense - Network Hardening", Hours: 4},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 4},
				{Name: "Network Defense - Defense in Depth", Hours: 4},
				{Name: "Network Defense - Honeypots and Honeynets", Hours: 4},
				{Name: "Network Defense - Exposure Minimization (Attack Surface and Vectors)", Hours: 4},
				{Name: "Network Defense - Network Access Control (Internal and External)", Hours: 4},
				{Name: "Network Defense - Perimeter Networks (DMZs) / Proxy Servers", Hours: 4},
				{Name: "Network Defense - Network Policy Development and Enforcement", Hours: 4},
				{Name: "Network Defense - Network Operational Procedures", Hours: 4},
				{Name: "Network Defense - Network Attacks (e.g., Session Hijacking, Man-in-the-Middle Attacks)", Hours: 6},
				{Name: "Network Implementations - Vulnerabilities and Exploit Examples", Hours: 4},
				{Name: "Network Services - Vulnerabilities and Exploit Examples", Hours: 4},
				{Name: "Distributed Systems Architecture - Vulnerabilities and Exploit Examples", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Digital Forensics - Reporting, Incident Response and Management", Hours: 4},
				{Name: "Digital Forensics - Mobile Forensics", Hours: 4},
				{Name: "Data Privacy - Overview", Hours: 6},
				{Name: "Privacy - Defining Privacy", Hours: 8},
				{Name: "Privacy - Privacy Rights", Hours: 6},
				{Name: "Privacy - Safeguarding Privacy", Hours: 4},
				{Name: "Privacy - Privacy Norms and Attitudes", Hours: 6},
				{Name: "Privacy - Privacy Breaches", Hours: 4},
				{Name: "Privacy - Privacy in Societies", Hours: 4},
				{Name: "Data Integrity and Authentication - Authentication Strength", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Attack Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Data Integrity", Hours: 4},
				{Name: "Information Storage Security - Database Security", Hours: 4},
				{Name: "Information Storage Security - Data Security Legislation", Hours: 4},
			},
		},
		{
			Name: "Software Security", Weight: 0.05,
			Topics: []topicCfg{
				{Name: "Analysis and Testing - Static and Dynamic Analysis", Hours: 6},
				{Name: "Analysis and Testing - Software Testing", Hours: 4},
				{Name: "Implementation - Input Validation and Representation Verification", Hours: 4},
				{Name: "Implementation - Correct Use of APIs", Hours: 4},
				{Name: "Implementation - Use of Security Features", Hours: 4},
				{Name: "Maintenance - Patching and Vulnerability Life Cycle", Hours: 4},
				{Name: "Component Testing - Security Testing", Hours: 4},
				{Name: "Legal Aspects - Vulnerability Disclosure", Hours: 4},
				{Name: "Social Engineering - Detection and Mitigation of Social Engineering Attacks", Hours: 4},
				{Name: "Awareness and Understanding - Awareness of Cyber Vulnerabilities and Threats", Hours: 4},
			},
		},
	}

	// 3) Insert KAs and Topics
	for _, ka := range cfg {
		kaID, err := findKA(ka.Name)
		if err != nil {
			return err
		}

		pak := model.ProfessionalRoleKnowledgeArea{
			ProfessionalRoleID: role.ID,
			KnowledgeAreaID:    kaID,
			PriorityWeight:     ka.Weight,
		}
		if err := dbCtx.DB().Create(&pak).Error; err != nil {
			return fmt.Errorf("cannot create role KA '%s': %w", ka.Name, err)
		}

		for _, t := range ka.Topics {
			topicID, err := findTopic(t.Name)
			if err != nil {
				return err
			}
			pat := model.ProfessionalRoleKnowledgeAreaTopic{
				ProfessionalRoleKnowledgeAreaID: pak.ID,
				TopicID:                         topicID,
				StudyHours:                      uint(t.Hours),
			}
			if err := dbCtx.DB().Create(&pat).Error; err != nil {
				return fmt.Errorf("cannot create KA topic '%s': %w", t.Name, err)
			}
		}
	}

	// 4) Assign Security Analyst role to all existing degree programs
	var degreePrograms []model.DegreeProgram
	if err := dbCtx.DB().Find(&degreePrograms).Error; err != nil {
		return fmt.Errorf("cannot find degree programs: %w", err)
	}

	for _, dp := range degreePrograms {
		// Check if the relationship already exists
		var existingRelation model.DegreeProgramProfessionalRole
		err := dbCtx.DB().Where("degree_program_id = ? AND professional_role_id = ?",
			dp.ID, role.ID).First(&existingRelation).Error

		if err != nil {
			// Relationship doesn't exist, create it
			relation := model.DegreeProgramProfessionalRole{
				DegreeProgramID:    dp.ID,
				ProfessionalRoleID: role.ID,
			}
			if err := dbCtx.DB().Create(&relation).Error; err != nil {
				return fmt.Errorf("cannot create relationship for degree program %d: %w", dp.ID, err)
			}
		}
	}

	return nil
}
