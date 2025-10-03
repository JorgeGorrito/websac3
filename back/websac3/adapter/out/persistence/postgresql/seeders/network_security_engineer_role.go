package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type networkSecurityEngineerRole struct{}

func NetworkSecurityEngineerRole() Seeder { return &networkSecurityEngineerRole{} }

func (s *networkSecurityEngineerRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Network Security Engineer
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Network Security Engineer").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Network Security Engineer"}
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

	// 2) Define KA expected for Network Security Engineer (weights sum to 1.0)
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
			Name: "Connection Security", Weight: 0.45,
			Topics: []topicCfg{
				{Name: "Network Architecture - General Concepts", Hours: 8},
				{Name: "Network Architecture - Common Architectures", Hours: 8},
				{Name: "Network Architecture - Forwarding", Hours: 6},
				{Name: "Network Architecture - Routing", Hours: 8},
				{Name: "Network Architecture - Switching/Bridging", Hours: 8},
				{Name: "Network Architecture - Emerging Trends", Hours: 6},
				{Name: "Network Architecture - Virtualization and Virtual Hypervisor Architectures", Hours: 8},
				{Name: "Network Implementations - IEEE 802/ISO Networks", Hours: 8},
				{Name: "Network Implementations - IETF and TCP/IP Networks", Hours: 10},
				{Name: "Network Implementations - Practical Integration and Bridging Protocols", Hours: 6},
				{Name: "Network Implementations - Vulnerabilities and Exploit Examples", Hours: 8},
				{Name: "Network Services - Concept of a Service", Hours: 6},
				{Name: "Network Services - Service Models (Client-Server, Peer-to-Peer)", Hours: 6},
				{Name: "Network Services - Service Protocol Concepts (IPC, API, IDL)", Hours: 6},
				{Name: "Network Services - Common Service Communication Architectures", Hours: 6},
				{Name: "Network Services - Service Virtualization", Hours: 8},
				{Name: "Network Services - Vulnerabilities and Exploit Examples", Hours: 6},
				{Name: "Network Defense - Network Hardening", Hours: 10},
				{Name: "Network Defense - IDS/IPS Implementation", Hours: 10},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 12},
				{Name: "Network Defense - Defense in Depth", Hours: 8},
				{Name: "Network Defense - Honeypots and Honeynets", Hours: 6},
				{Name: "Network Defense - Network Monitoring", Hours: 8},
				{Name: "Network Defense - Network Traffic Analysis", Hours: 8},
				{Name: "Network Defense - Exposure Minimization (Attack Surface and Vectors)", Hours: 6},
				{Name: "Network Defense - Network Access Control (Internal and External)", Hours: 8},
				{Name: "Network Defense - Perimeter Networks (DMZs) / Proxy Servers", Hours: 8},
				{Name: "Network Defense - Network Policy Development and Enforcement", Hours: 6},
				{Name: "Network Defense - Network Operational Procedures", Hours: 6},
				{Name: "Network Defense - Network Attacks (e.g., Session Hijacking, Man-in-the-Middle Attacks)", Hours: 8},
				{Name: "Network Defense - Threat Hunting and Machine Learning", Hours: 6},
			},
		},
		{
			Name: "Connection Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Cryptography - Basic Concepts", Hours: 6},
				{Name: "Cryptography - Advanced Concepts", Hours: 6},
				{Name: "Cryptography - Mathematical Foundations", Hours: 4},
				{Name: "Cryptography - Historical Ciphers", Hours: 4},
				{Name: "Cryptography - Symmetric Ciphers (Private Key)", Hours: 6},
				{Name: "Cryptography - Asymmetric Ciphers (Public Key)", Hours: 6},
				{Name: "Secure Communication Protocols - Application and Transport Layer Protocols", Hours: 8},
				{Name: "Secure Communication Protocols - TLS Attacks", Hours: 6},
				{Name: "Secure Communication Protocols - Network/Internet Layer", Hours: 8},
				{Name: "Secure Communication Protocols - Privacy Preservation Protocols", Hours: 6},
				{Name: "Secure Communication Protocols - Data Link Layer", Hours: 6},
				{Name: "Cryptanalysis - Classical Attacks", Hours: 4},
				{Name: "Cryptanalysis - Side Channel Attacks", Hours: 4},
				{Name: "Cryptanalysis - Attacks on Private Key Ciphers", Hours: 4},
				{Name: "Cryptanalysis - Attacks on Public Key Ciphers", Hours: 4},
				{Name: "Cryptanalysis - Algorithms for the Discrete Logarithm Problem", Hours: 4},
				{Name: "Cryptanalysis - Attacks on RSA", Hours: 4},
			},
		},
		{
			Name: "System Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "System Administration - Network Administration", Hours: 10},
				{Name: "System Administration - Operating Systems Administration", Hours: 6},
				{Name: "System Administration - System Hardening", Hours: 8},
				{Name: "System Administration - Availability", Hours: 6},
				{Name: "Systems Management - Policy Models", Hours: 4},
				{Name: "Systems Management - Policy Composition", Hours: 4},
				{Name: "Systems Management - Use of Automation", Hours: 4},
				{Name: "Systems Management - Patching and Vulnerability Life Cycle", Hours: 6},
				{Name: "Systems Management - Operation", Hours: 6},
				{Name: "Systems Management - Deployment and Retirement", Hours: 6},
				{Name: "Systems Management - Insider Threat", Hours: 4},
				{Name: "Systems Management - Documentation", Hours: 4},
				{Name: "Systems Management - Systems and Procedures", Hours: 4},
				{Name: "System Access - Authentication Methods", Hours: 4},
				{Name: "System Access - Identity", Hours: 4},
				{Name: "System Control - Access Control", Hours: 6},
				{Name: "System Control - Authorization Models", Hours: 6},
				{Name: "System Control - Intrusion Detection", Hours: 6},
				{Name: "System Control - Attacks", Hours: 6},
				{Name: "System Control - Defenses", Hours: 6},
				{Name: "System Control - Auditing", Hours: 4},
				{Name: "System Control - Malware", Hours: 4},
				{Name: "System Control - Vulnerability Models", Hours: 6},
				{Name: "System Control - Penetration Testing", Hours: 4},
				{Name: "System Control - Forensics", Hours: 4},
				{Name: "System Control - Recovery and Resilience", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.10,
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
			Name: "Data Security", Weight: 0.05,
			Topics: []topicCfg{
				{Name: "Information Storage Security - Database Security", Hours: 4},
				{Name: "Information Storage Security - Disk and File Encryption", Hours: 4},
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

	// 4) Assign Network Security Engineer role to all existing degree programs
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
