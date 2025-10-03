package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type securityArchitectRole struct{}

func SecurityArchitectRole() Seeder { return &securityArchitectRole{} }

func (s *securityArchitectRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Security Architect
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Security Architect").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Security Architect"}
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

	// 2) Define KA expected for Security Architect (weights sum to 1.0)
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
				{Name: "Systems Management - Policy Models", Hours: 6},
				{Name: "Systems Management - Policy Composition", Hours: 6},
				{Name: "Systems Management - Use of Automation", Hours: 6},
				{Name: "Systems Management - Patching and Vulnerability Life Cycle", Hours: 6},
				{Name: "Systems Management - Operation", Hours: 6},
				{Name: "Systems Management - Deployment and Retirement", Hours: 4},
				{Name: "Systems Management - Insider Threat", Hours: 6},
				{Name: "Systems Management - Documentation", Hours: 4},
				{Name: "Systems Management - Systems and Procedures", Hours: 4},
				{Name: "System Access - Authentication Methods", Hours: 6},
				{Name: "System Access - Identity", Hours: 6},
				{Name: "System Control - Access Control", Hours: 6},
				{Name: "System Control - Authorization Models", Hours: 6},
				{Name: "System Control - Intrusion Detection", Hours: 6},
			},
		},
		{
			Name: "Software Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "Implementation - Encapsulation of Structures and Modules", Hours: 8},
				{Name: "Implementation - Consideration of the Environment", Hours: 8},
				{Name: "Analysis and Testing - Static and Dynamic Analysis", Hours: 8},
				{Name: "Analysis and Testing - Unit Testing", Hours: 6},
				{Name: "Analysis and Testing - Integration Testing", Hours: 6},
				{Name: "Analysis and Testing - Software Testing", Hours: 6},
				{Name: "Maintenance - Configuration", Hours: 4},
				{Name: "Maintenance - Patching and Vulnerability Life Cycle", Hours: 6},
				{Name: "Maintenance - Environment Verification", Hours: 6},
				{Name: "Maintenance - DevOps", Hours: 6},
			},
		},
		{
			Name: "Connection Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Network Architecture - General Concepts", Hours: 6},
				{Name: "Network Architecture - Common Architectures", Hours: 6},
				{Name: "Network Architecture - Forwarding", Hours: 4},
				{Name: "Network Architecture - Routing", Hours: 6},
				{Name: "Network Architecture - Switching/Bridging", Hours: 4},
				{Name: "Network Architecture - Emerging Trends", Hours: 4},
				{Name: "Network Architecture - Virtualization and Virtual Hypervisor Architectures", Hours: 6},
				{Name: "Network Implementations - IEEE 802/ISO Networks", Hours: 4},
				{Name: "Network Implementations - IETF and TCP/IP Networks", Hours: 6},
				{Name: "Network Implementations - Practical Integration and Bridging Protocols", Hours: 4},
				{Name: "Network Implementations - Vulnerabilities and Exploit Examples", Hours: 6},
				{Name: "Network Services - Concept of a Service", Hours: 4},
				{Name: "Network Services - Service Models (Client-Server, Peer-to-Peer)", Hours: 4},
				{Name: "Network Services - Service Protocol Concepts (IPC, API, IDL)", Hours: 4},
				{Name: "Network Services - Common Service Communication Architectures", Hours: 4},
				{Name: "Network Services - Service Virtualization", Hours: 4},
				{Name: "Network Services - Vulnerabilities and Exploit Examples", Hours: 4},
				{Name: "Network Defense - Network Hardening", Hours: 6},
				{Name: "Network Defense - IDS/IPS Implementation", Hours: 6},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 8},
				{Name: "Network Defense - Defense in Depth", Hours: 6},
				{Name: "Network Defense - Honeypots and Honeynets", Hours: 4},
				{Name: "Network Defense - Network Monitoring", Hours: 6},
				{Name: "Network Defense - Network Traffic Analysis", Hours: 6},
				{Name: "Network Defense - Exposure Minimization (Attack Surface and Vectors)", Hours: 6},
				{Name: "Network Defense - Network Access Control (Internal and External)", Hours: 6},
				{Name: "Network Defense - Perimeter Networks (DMZs) / Proxy Servers", Hours: 6},
				{Name: "Network Defense - Network Policy Development and Enforcement", Hours: 6},
				{Name: "Network Defense - Network Operational Procedures", Hours: 4},
				{Name: "Network Defense - Network Attacks (e.g., Session Hijacking, Man-in-the-Middle Attacks)", Hours: 6},
				{Name: "Network Defense - Threat Hunting and Machine Learning", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Access Control - Secure Architecture Design", Hours: 8},
				{Name: "Access Control - Physical Data Security", Hours: 4},
				{Name: "Access Control - Logical Data Access Control", Hours: 6},
				{Name: "Access Control - Data Leakage Prevention Techniques", Hours: 6},
				{Name: "Data Integrity and Authentication - Authentication Strength", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Attack Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Data Integrity", Hours: 4},
				{Name: "Information Storage Security - Disk and File Encryption", Hours: 6},
				{Name: "Information Storage Security - Data Erasure", Hours: 4},
				{Name: "Information Storage Security - Data Masking", Hours: 4},
				{Name: "Information Storage Security - Database Security", Hours: 8},
				{Name: "Information Storage Security - Data Security Legislation", Hours: 4},
				{Name: "Data Privacy - Overview", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "System Administration - Operating Systems Administration", Hours: 6},
				{Name: "System Administration - Database System Administration", Hours: 6},
				{Name: "System Administration - Network Administration", Hours: 6},
				{Name: "System Administration - Cloud Administration", Hours: 6},
				{Name: "System Administration - Cyber-Physical Systems Administration", Hours: 4},
				{Name: "System Administration - System Hardening", Hours: 6},
				{Name: "System Administration - Availability", Hours: 4},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 4},
				{Name: "Identity Management - Identification and Authentication of People and Devices", Hours: 6},
				{Name: "Identity Management - Physical and Logical Asset Control", Hours: 4},
				{Name: "Identity Management - Identity as a Service (IDaaS)", Hours: 4},
				{Name: "Identity Management - Third-Party Identity Services", Hours: 4},
				{Name: "Identity Management - Access Control Attacks and Mitigations", Hours: 4},
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

	// 4) Assign Security Architect role to all existing degree programs
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
