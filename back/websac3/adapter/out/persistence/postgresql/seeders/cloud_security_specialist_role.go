package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type cloudSecuritySpecialistRole struct{}

func CloudSecuritySpecialistRole() Seeder { return &cloudSecuritySpecialistRole{} }

func (s *cloudSecuritySpecialistRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Cloud Security Specialist
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Cloud Security Specialist").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Cloud Security Specialist"}
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

	// 2) Define KA expected for Cloud Security Specialist (weights sum to 1.0)
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
			Name: "System Security", Weight: 0.35,
			Topics: []topicCfg{
				{Name: "System Administration - Cloud Administration", Hours: 12},
				{Name: "System Administration - Operating Systems Administration", Hours: 8},
				{Name: "System Administration - Database System Administration", Hours: 8},
				{Name: "System Administration - Network Administration", Hours: 8},
				{Name: "System Administration - Cyber-Physical Systems Administration", Hours: 6},
				{Name: "System Administration - System Hardening", Hours: 10},
				{Name: "System Administration - Availability", Hours: 8},
				{Name: "Systems Management - Policy Models", Hours: 6},
				{Name: "Systems Management - Policy Composition", Hours: 6},
				{Name: "Systems Management - Use of Automation", Hours: 6},
				{Name: "Systems Management - Patching and Vulnerability Life Cycle", Hours: 6},
				{Name: "Systems Management - Operation", Hours: 6},
				{Name: "Systems Management - Deployment and Retirement", Hours: 6},
				{Name: "Systems Management - Insider Threat", Hours: 6},
				{Name: "Systems Management - Documentation", Hours: 4},
				{Name: "Systems Management - Systems and Procedures", Hours: 4},
				{Name: "System Access - Authentication Methods", Hours: 4},
				{Name: "System Access - Identity", Hours: 4},
				{Name: "System Control - Access Control", Hours: 6},
				{Name: "System Control - Authorization Models", Hours: 6},
				{Name: "System Control - Intrusion Detection", Hours: 6},
				{Name: "System Control - Attacks", Hours: 6},
				{Name: "System Control - Defenses", Hours: 6},
				{Name: "System Control - Auditing", Hours: 6},
				{Name: "System Control - Malware", Hours: 4},
				{Name: "System Control - Vulnerability Models", Hours: 6},
				{Name: "System Control - Penetration Testing", Hours: 4},
				{Name: "System Control - Forensics", Hours: 4},
				{Name: "System Control - Recovery and Resilience", Hours: 6},
				{Name: "System Testing - Requirements Validation", Hours: 4},
				{Name: "System Testing - Component Composition Validation", Hours: 4},
				{Name: "System Testing - Unit Testing versus System Testing", Hours: 4},
				{Name: "System Testing - Formal Verification of Systems", Hours: 4},
			},
		},
		{
			Name: "Connection Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "Distributed Systems Architecture - General Concepts", Hours: 6},
				{Name: "Distributed Systems Architecture - World Wide Web", Hours: 4},
				{Name: "Distributed Systems Architecture - Internet", Hours: 6},
				{Name: "Distributed Systems Architecture - Protocols and Layers", Hours: 6},
				{Name: "Distributed Systems Architecture - High Performance Computing (Supercomputers)", Hours: 4},
				{Name: "Distributed Systems Architecture - Hypervisors and Cloud Computing Implementations", Hours: 10},
				{Name: "Distributed Systems Architecture - Vulnerabilities and Exploit Examples", Hours: 6},
				{Name: "Network Architecture - General Concepts", Hours: 4},
				{Name: "Network Architecture - Common Architectures", Hours: 4},
				{Name: "Network Architecture - Forwarding", Hours: 4},
				{Name: "Network Architecture - Routing", Hours: 4},
				{Name: "Network Architecture - Switching/Bridging", Hours: 4},
				{Name: "Network Architecture - Emerging Trends", Hours: 4},
				{Name: "Network Architecture - Virtualization and Virtual Hypervisor Architectures", Hours: 8},
				{Name: "Network Implementations - IEEE 802/ISO Networks", Hours: 4},
				{Name: "Network Implementations - IETF and TCP/IP Networks", Hours: 6},
				{Name: "Network Implementations - Practical Integration and Bridging Protocols", Hours: 4},
				{Name: "Network Implementations - Vulnerabilities and Exploit Examples", Hours: 4},
				{Name: "Network Services - Concept of a Service", Hours: 4},
				{Name: "Network Services - Service Models (Client-Server, Peer-to-Peer)", Hours: 4},
				{Name: "Network Services - Service Protocol Concepts (IPC, API, IDL)", Hours: 4},
				{Name: "Network Services - Common Service Communication Architectures", Hours: 4},
				{Name: "Network Services - Service Virtualization", Hours: 6},
				{Name: "Network Services - Vulnerabilities and Exploit Examples", Hours: 4},
				{Name: "Network Defense - Network Hardening", Hours: 6},
				{Name: "Network Defense - IDS/IPS Implementation", Hours: 4},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 6},
				{Name: "Network Defense - Defense in Depth", Hours: 4},
				{Name: "Network Defense - Honeypots and Honeynets", Hours: 4},
				{Name: "Network Defense - Network Monitoring", Hours: 4},
				{Name: "Network Defense - Network Traffic Analysis", Hours: 4},
				{Name: "Network Defense - Exposure Minimization (Attack Surface and Vectors)", Hours: 4},
				{Name: "Network Defense - Network Access Control (Internal and External)", Hours: 4},
				{Name: "Network Defense - Perimeter Networks (DMZs) / Proxy Servers", Hours: 4},
				{Name: "Network Defense - Network Policy Development and Enforcement", Hours: 4},
				{Name: "Network Defense - Network Operational Procedures", Hours: 4},
				{Name: "Network Defense - Network Attacks (e.g., Session Hijacking, Man-in-the-Middle Attacks)", Hours: 4},
				{Name: "Network Defense - Threat Hunting and Machine Learning", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Information Storage Security - Database Security", Hours: 8},
				{Name: "Information Storage Security - Disk and File Encryption", Hours: 6},
				{Name: "Information Storage Security - Data Erasure", Hours: 4},
				{Name: "Information Storage Security - Data Masking", Hours: 4},
				{Name: "Information Storage Security - Data Security Legislation", Hours: 4},
				{Name: "Data Integrity and Authentication - Authentication Strength", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Attack Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Data Integrity", Hours: 4},
				{Name: "Access Control - Physical Data Security", Hours: 4},
				{Name: "Access Control - Logical Data Access Control", Hours: 4},
				{Name: "Access Control - Secure Architecture Design", Hours: 4},
				{Name: "Access Control - Data Leakage Prevention Techniques", Hours: 4},
				{Name: "Data Privacy - Overview", Hours: 4},
				{Name: "Digital Forensics - Introduction", Hours: 4},
				{Name: "Digital Forensics - Tools", Hours: 4},
				{Name: "Digital Forensics - Data Analysis Techniques", Hours: 4},
				{Name: "Digital Forensics - Investigation Process", Hours: 4},
				{Name: "Digital Forensics - Evidence Acquisition and Preservation", Hours: 4},
				{Name: "Digital Forensics - Evidence Analysis", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Security Program Management - Project Management", Hours: 4},
				{Name: "Security Program Management - Resource Management", Hours: 4},
				{Name: "Security Program Management - Security Metrics", Hours: 4},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 4},
				{Name: "Analytical Tools - Performance Measurement (Metrics)", Hours: 4},
				{Name: "Analytical Tools - Data Analytics", Hours: 4},
				{Name: "Analytical Tools - Security Intelligence", Hours: 4},
				{Name: "Security Operations - Security Convergence", Hours: 4},
				{Name: "Security Operations - Global Security Operations Centers (GSOCs)", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Disaster Recovery", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Business Continuity", Hours: 4},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 4},
				{Name: "Risk Management - Risk Identification", Hours: 4},
				{Name: "Risk Management - Risk Assessment and Analysis", Hours: 4},
				{Name: "Risk Management - Insider Threats", Hours: 4},
				{Name: "Risk Management - Risk Control", Hours: 4},
				{Name: "Personnel Security - Third-Party Security", Hours: 4},
				{Name: "Personnel Security - Secure Hiring Practices", Hours: 4},
				{Name: "Personnel Security - Secure Termination Practices", Hours: 4},
			},
		},
		{
			Name: "Software Security", Weight: 0.05,
			Topics: []topicCfg{
				{Name: "Analysis and Testing - Static and Dynamic Analysis", Hours: 4},
				{Name: "Analysis and Testing - Software Testing", Hours: 4},
				{Name: "Implementation - Input Validation and Representation Verification", Hours: 4},
				{Name: "Implementation - Correct Use of APIs", Hours: 4},
				{Name: "Implementation - Use of Security Features", Hours: 4},
				{Name: "Maintenance - Configuration", Hours: 4},
				{Name: "Maintenance - Patching and Vulnerability Life Cycle", Hours: 4},
				{Name: "Maintenance - Environment Verification", Hours: 4},
				{Name: "Maintenance - DevOps", Hours: 4},
				{Name: "Legal Aspects - Vulnerability Disclosure", Hours: 4},
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

	// 4) Assign Cloud Security Specialist role to all existing degree programs
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
