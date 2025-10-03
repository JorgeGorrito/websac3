package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type iamSpecialistRole struct{}

func IAMSpecialistRole() Seeder { return &iamSpecialistRole{} }

func (s *iamSpecialistRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Identity and Access Management (IAM) Specialist
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Identity and Access Management (IAM) Specialist").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Identity and Access Management (IAM) Specialist"}
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

	// 2) Define KA expected for IAM Specialist (weights sum to 1.0)
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
			Name: "Data Security", Weight: 0.40,
			Topics: []topicCfg{
				{Name: "Identity Management - Identification and Authentication of People and Devices", Hours: 12},
				{Name: "Identity Management - Physical and Logical Asset Control", Hours: 10},
				{Name: "Identity Management - Identity as a Service (IDaaS)", Hours: 10},
				{Name: "Identity Management - Third-Party Identity Services", Hours: 8},
				{Name: "Identity Management - Access Control Attacks and Mitigations", Hours: 8},
				{Name: "Data Integrity and Authentication - Authentication Strength", Hours: 8},
				{Name: "Data Integrity and Authentication - Password Attack Techniques", Hours: 6},
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 8},
				{Name: "Data Integrity and Authentication - Data Integrity", Hours: 6},
				{Name: "Access Control - Physical Data Security", Hours: 6},
				{Name: "Access Control - Logical Data Access Control", Hours: 8},
				{Name: "Access Control - Secure Architecture Design", Hours: 6},
				{Name: "Access Control - Data Leakage Prevention Techniques", Hours: 6},
				{Name: "Information Storage Security - Database Security", Hours: 6},
				{Name: "Information Storage Security - Data Security Legislation", Hours: 4},
				{Name: "Data Privacy - Overview", Hours: 4},
			},
		},
		{
			Name: "System Security", Weight: 0.30,
			Topics: []topicCfg{
				{Name: "System Access - Authentication Methods", Hours: 10},
				{Name: "System Access - Identity", Hours: 8},
				{Name: "System Control - Access Control", Hours: 10},
				{Name: "System Control - Authorization Models", Hours: 8},
				{Name: "System Control - Intrusion Detection", Hours: 6},
				{Name: "System Control - Attacks", Hours: 6},
				{Name: "System Control - Defenses", Hours: 6},
				{Name: "System Control - Auditing", Hours: 8},
				{Name: "System Control - Vulnerability Models", Hours: 4},
				{Name: "System Control - Recovery and Resilience", Hours: 4},
				{Name: "Systems Management - Policy Models", Hours: 6},
				{Name: "Systems Management - Policy Composition", Hours: 6},
				{Name: "Systems Management - Use of Automation", Hours: 4},
				{Name: "Systems Management - Insider Threat", Hours: 6},
				{Name: "Systems Management - Documentation", Hours: 4},
				{Name: "Systems Management - Systems and Procedures", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Security Program Management - Project Management", Hours: 6},
				{Name: "Security Program Management - Resource Management", Hours: 6},
				{Name: "Security Program Management - Security Metrics", Hours: 6},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 4},
				{Name: "Personnel Security - Third-Party Security", Hours: 6},
				{Name: "Personnel Security - Secure Hiring Practices", Hours: 4},
				{Name: "Personnel Security - Secure Termination Practices", Hours: 4},
				{Name: "Personnel Security - Security in Review Processes", Hours: 4},
				{Name: "Personnel Security - Special Issues of Employee Personal Information Privacy", Hours: 4},
				{Name: "Analytical Tools - Performance Measurement (Metrics)", Hours: 4},
				{Name: "Analytical Tools - Data Analytics", Hours: 4},
				{Name: "Analytical Tools - Security Intelligence", Hours: 4},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 4},
				{Name: "Risk Management - Risk Identification", Hours: 4},
				{Name: "Risk Management - Risk Assessment and Analysis", Hours: 4},
				{Name: "Risk Management - Insider Threats", Hours: 6},
				{Name: "Risk Management - Risk Control", Hours: 4},
			},
		},
		{
			Name: "Connection Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "Network Defense - Network Access Control (Internal and External)", Hours: 6},
				{Name: "Network Defense - Network Monitoring", Hours: 4},
				{Name: "Network Defense - Network Traffic Analysis", Hours: 4},
				{Name: "Network Defense - IDS/IPS Implementation", Hours: 4},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 4},
				{Name: "Network Defense - Defense in Depth", Hours: 4},
				{Name: "Network Defense - Network Attacks (e.g., Session Hijacking, Man-in-the-Middle Attacks)", Hours: 4},
				{Name: "Network Services - Concept of a Service", Hours: 4},
				{Name: "Network Services - Service Models (Client-Server, Peer-to-Peer)", Hours: 4},
				{Name: "Network Services - Service Protocol Concepts (IPC, API, IDL)", Hours: 4},
				{Name: "Network Services - Service Virtualization", Hours: 4},
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

	// 4) Assign IAM Specialist role to all existing degree programs
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
