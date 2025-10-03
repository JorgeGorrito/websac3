package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type securityEngineerRole struct{}

func SecurityEngineerRole() Seeder { return &securityEngineerRole{} }

func (s *securityEngineerRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Security Engineer
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Security Engineer").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Security Engineer"}
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

	// 2) Define KA expected for Security Engineer (weights sum to 1.0)
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
			Name: "System Security", Weight: 0.30,
			Topics: []topicCfg{
				{Name: "System Control - Intrusion Detection", Hours: 10},
				{Name: "System Control - Attacks", Hours: 8},
				{Name: "System Control - Defenses", Hours: 10},
				{Name: "System Control - Auditing", Hours: 6},
				{Name: "System Control - Malware", Hours: 8},
				{Name: "System Control - Vulnerability Models", Hours: 8},
				{Name: "System Control - Penetration Testing", Hours: 8},
				{Name: "System Control - Forensics", Hours: 6},
				{Name: "System Control - Recovery and Resilience", Hours: 6},
				{Name: "Systems Management - Use of Automation", Hours: 6},
				{Name: "Systems Management - Patching and Vulnerability Life Cycle", Hours: 8},
				{Name: "Systems Management - Operation", Hours: 6},
				{Name: "Systems Management - Insider Threat", Hours: 6},
				{Name: "Systems Management - Documentation", Hours: 4},
				{Name: "Systems Management - Systems and Procedures", Hours: 4},
				{Name: "System Testing - Requirements Validation", Hours: 4},
				{Name: "System Testing - Component Composition Validation", Hours: 4},
				{Name: "System Testing - Unit Testing versus System Testing", Hours: 4},
				{Name: "System Testing - Formal Verification of Systems", Hours: 4},
			},
		},
		{
			Name: "Connection Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "Network Defense - Network Hardening", Hours: 8},
				{Name: "Network Defense - IDS/IPS Implementation", Hours: 10},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 10},
				{Name: "Network Defense - Defense in Depth", Hours: 6},
				{Name: "Network Defense - Honeypots and Honeynets", Hours: 4},
				{Name: "Network Defense - Network Monitoring", Hours: 8},
				{Name: "Network Defense - Network Traffic Analysis", Hours: 8},
				{Name: "Network Defense - Exposure Minimization (Attack Surface and Vectors)", Hours: 6},
				{Name: "Network Defense - Network Access Control (Internal and External)", Hours: 6},
				{Name: "Network Defense - Perimeter Networks (DMZs) / Proxy Servers", Hours: 6},
				{Name: "Network Defense - Network Policy Development and Enforcement", Hours: 6},
				{Name: "Network Defense - Network Operational Procedures", Hours: 4},
				{Name: "Network Defense - Network Attacks (e.g., Session Hijacking, Man-in-the-Middle Attacks)", Hours: 8},
				{Name: "Network Defense - Threat Hunting and Machine Learning", Hours: 6},
				{Name: "Network Implementations - Vulnerabilities and Exploit Examples", Hours: 6},
				{Name: "Network Services - Vulnerabilities and Exploit Examples", Hours: 4},
				{Name: "Secure Communication Protocols - TLS Attacks", Hours: 6},
			},
		},
		{
			Name: "Software Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Implementation - Input Validation and Representation Verification", Hours: 8},
				{Name: "Implementation - Correct Use of APIs", Hours: 6},
				{Name: "Implementation - Use of Security Features", Hours: 6},
				{Name: "Implementation - Verification of Time and State Relationships", Hours: 4},
				{Name: "Implementation - Proper Handling of Exceptions and Errors", Hours: 4},
				{Name: "Implementation - Robust Programming", Hours: 6},
				{Name: "Implementation - Encapsulation of Structures and Modules", Hours: 4},
				{Name: "Implementation - Consideration of the Environment", Hours: 4},
				{Name: "Analysis and Testing - Static and Dynamic Analysis", Hours: 8},
				{Name: "Analysis and Testing - Unit Testing", Hours: 4},
				{Name: "Analysis and Testing - Integration Testing", Hours: 4},
				{Name: "Analysis and Testing - Software Testing", Hours: 4},
				{Name: "Maintenance - Configuration", Hours: 4},
				{Name: "Maintenance - Patching and Vulnerability Life Cycle", Hours: 6},
				{Name: "Maintenance - Environment Verification", Hours: 4},
				{Name: "Maintenance - DevOps", Hours: 6},
				{Name: "Maintenance - Retirement/Decommissioning", Hours: 4},
				{Name: "Component Testing - Security Testing", Hours: 6},
				{Name: "Legal Aspects - Vulnerability Disclosure", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Security Operations - Security Convergence", Hours: 6},
				{Name: "Security Operations - Global Security Operations Centers (GSOCs)", Hours: 6},
				{Name: "System Administration - Operating Systems Administration", Hours: 6},
				{Name: "System Administration - Database System Administration", Hours: 4},
				{Name: "System Administration - Network Administration", Hours: 6},
				{Name: "System Administration - Cloud Administration", Hours: 6},
				{Name: "System Administration - Cyber-Physical Systems Administration", Hours: 4},
				{Name: "System Administration - System Hardening", Hours: 8},
				{Name: "System Administration - Availability", Hours: 4},
				{Name: "Security Program Management - Security Metrics", Hours: 6},
				{Name: "Security Operations - Security Convergence", Hours: 6},
				{Name: "Security Operations - Global Security Operations Centers (GSOCs)", Hours: 6},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 6},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 6},
				{Name: "Security Program Management - Project Management", Hours: 4},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "Digital Forensics - Reporting, Incident Response and Management", Hours: 4},
				{Name: "Digital Forensics - Mobile Forensics", Hours: 4},
				{Name: "Data Privacy - Overview", Hours: 4},
				{Name: "Privacy - Defining Privacy", Hours: 4},
				{Name: "Privacy - Privacy Rights", Hours: 4},
				{Name: "Privacy - Safeguarding Privacy", Hours: 4},
				{Name: "Privacy - Privacy Norms and Attitudes", Hours: 4},
				{Name: "Privacy - Privacy Breaches", Hours: 4},
				{Name: "Privacy - Privacy in Societies", Hours: 4},
				{Name: "Data Integrity and Authentication - Authentication Strength", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Attack Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Data Integrity", Hours: 4},
				{Name: "Information Storage Security - Disk and File Encryption", Hours: 4},
				{Name: "Information Storage Security - Data Erasure", Hours: 4},
				{Name: "Information Storage Security - Data Masking", Hours: 4},
				{Name: "Information Storage Security - Database Security", Hours: 6},
				{Name: "Information Storage Security - Data Security Legislation", Hours: 4},
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

	// 4) Assign Security Engineer role to all existing degree programs
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
