package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type cybersecurityTrainerAwarenessSpecialistRole struct{}

func CybersecurityTrainerAwarenessSpecialistRole() Seeder {
	return &cybersecurityTrainerAwarenessSpecialistRole{}
}

func (s *cybersecurityTrainerAwarenessSpecialistRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Cybersecurity Trainer / Awareness Specialist
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Cybersecurity Trainer / Awareness Specialist").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Cybersecurity Trainer / Awareness Specialist"}
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

	// 2) Define KA expected for Cybersecurity Trainer / Awareness Specialist (weights sum to 1.0)
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
			Name: "Human Security", Weight: 0.45,
			Topics: []topicCfg{
				{Name: "Awareness and Understanding - Risk Perception and Communication", Hours: 12},
				{Name: "Awareness and Understanding - Cyber Hygiene", Hours: 10},
				{Name: "Awareness and Understanding - Cybersecurity User Education", Hours: 12},
				{Name: "Awareness and Understanding - Awareness of Cyber Vulnerabilities and Threats", Hours: 10},
				{Name: "Personal Compliance with Cybersecurity Rules/Policies/Ethical Norms - System Misuse and User Misbehavior", Hours: 8},
				{Name: "Personal Compliance with Cybersecurity Rules/Policies/Ethical Norms - Enforcement and Rules of Behavior", Hours: 8},
				{Name: "Personal Compliance with Cybersecurity Rules/Policies/Ethical Norms - Appropriate Behavior under Uncertainty", Hours: 6},
				{Name: "Usable Security and Privacy - Usability and User Experience", Hours: 8},
				{Name: "Usable Security and Privacy - Human Security Factors", Hours: 6},
				{Name: "Usable Security and Privacy - Policy Awareness and Understanding", Hours: 8},
				{Name: "Usable Security and Privacy - Privacy Policy", Hours: 6},
				{Name: "Usable Security and Privacy - Design Guidance and Implications", Hours: 6},
				{Name: "Social and Personal Privacy - Privacy and Security in Social Networks", Hours: 6},
				{Name: "Identity Management - Physical and Logical Asset Control", Hours: 4},
				{Name: "Identity Management - Access Control Attacks and Mitigations", Hours: 4},
				{Name: "Social Engineering - Types of Social Engineering Attacks", Hours: 6},
				{Name: "Social Engineering - Psychology of Social Engineering Attacks", Hours: 6},
				{Name: "Social Engineering - Detection and Mitigation of Social Engineering Attacks", Hours: 6},
				{Name: "Social Engineering - Deception of Users", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "Personnel Security - Security Awareness, Training, and Education", Hours: 12},
				{Name: "Personnel Security - Secure Hiring Practices", Hours: 6},
				{Name: "Personnel Security - Secure Termination Practices", Hours: 6},
				{Name: "Personnel Security - Third-Party Security", Hours: 6},
				{Name: "Personnel Security - Security in Review Processes", Hours: 4},
				{Name: "Personnel Security - Special Issues of Employee Personal Information Privacy", Hours: 6},
				{Name: "Security Program Management - Project Management", Hours: 6},
				{Name: "Security Program Management - Resource Management", Hours: 6},
				{Name: "Security Program Management - Security Metrics", Hours: 8},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 6},
				{Name: "Analytical Tools - Performance Measurement (Metrics)", Hours: 6},
				{Name: "Analytical Tools - Data Analytics", Hours: 6},
				{Name: "Analytical Tools - Security Intelligence", Hours: 4},
				{Name: "Security Operations - Security Convergence", Hours: 4},
				{Name: "Security Operations - Global Security Operations Centers (GSOCs)", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Disaster Recovery", Hours: 4},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Business Continuity", Hours: 4},
				{Name: "Cybersecurity Planning - Strategic Planning", Hours: 4},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 4},
				{Name: "Risk Management - Risk Identification", Hours: 6},
				{Name: "Risk Management - Risk Assessment and Analysis", Hours: 6},
				{Name: "Risk Management - Insider Threats", Hours: 4},
				{Name: "Risk Management - Risk Measurement and Evaluation Models and Methodologies", Hours: 4},
				{Name: "Risk Management - Risk Control", Hours: 4},
				{Name: "Security Governance and Policy - Organizational Context", Hours: 4},
				{Name: "Security Governance and Policy - Privacy", Hours: 4},
				{Name: "Security Governance and Policy - Law, Ethics, and Compliance", Hours: 4},
				{Name: "Security Governance and Policy - Security Governance", Hours: 4},
				{Name: "Security Governance and Policy - Executive and Board-Level Communication", Hours: 6},
				{Name: "Security Governance and Policy - Management Policy", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Data Privacy - Overview", Hours: 8},
				{Name: "Personal Data Privacy and Security - Sensitive Personal Data (SPD)", Hours: 8},
				{Name: "Personal Data Privacy and Security - Personal Tracking and Digital Footprint", Hours: 6},
				{Name: "Privacy - Privacy Rights", Hours: 6},
				{Name: "Information Storage Security - Database Security", Hours: 4},
				{Name: "Information Storage Security - Disk and File Encryption", Hours: 4},
				{Name: "Information Storage Security - Data Erasure", Hours: 4},
				{Name: "Information Storage Security - Data Masking", Hours: 4},
				{Name: "Information Storage Security - Data Security Legislation", Hours: 6},
				{Name: "Data Integrity and Authentication - Authentication Strength", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Attack Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 4},
				{Name: "Data Integrity and Authentication - Data Integrity", Hours: 4},
				{Name: "Access Control - Physical Data Security", Hours: 4},
				{Name: "Access Control - Logical Data Access Control", Hours: 4},
				{Name: "Access Control - Secure Architecture Design", Hours: 4},
				{Name: "Access Control - Data Leakage Prevention Techniques", Hours: 4},
			},
		},
		{
			Name: "System Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "System Control - Access Control", Hours: 4},
				{Name: "System Control - Authorization Models", Hours: 4},
				{Name: "System Control - Intrusion Detection", Hours: 4},
				{Name: "System Control - Attacks", Hours: 6},
				{Name: "System Control - Defenses", Hours: 6},
				{Name: "System Control - Auditing", Hours: 4},
				{Name: "System Control - Malware", Hours: 6},
				{Name: "System Control - Vulnerability Models", Hours: 4},
				{Name: "System Control - Penetration Testing", Hours: 4},
				{Name: "System Control - Forensics", Hours: 4},
				{Name: "System Control - Recovery and Resilience", Hours: 4},
				{Name: "Systems Management - Policy Models", Hours: 4},
				{Name: "Systems Management - Policy Composition", Hours: 4},
				{Name: "Systems Management - Use of Automation", Hours: 4},
				{Name: "Systems Management - Patching and Vulnerability Life Cycle", Hours: 4},
				{Name: "Systems Management - Operation", Hours: 4},
				{Name: "Systems Management - Deployment and Retirement", Hours: 4},
				{Name: "Systems Management - Insider Threat", Hours: 4},
				{Name: "Systems Management - Documentation", Hours: 4},
				{Name: "Systems Management - Systems and Procedures", Hours: 4},
				{Name: "System Access - Authentication Methods", Hours: 4},
				{Name: "System Access - Identity", Hours: 4},
			},
		},
		{
			Name: "Human Security", Weight: 0.05,
			Topics: []topicCfg{
				{Name: "Cyberethics - Defining Ethics", Hours: 4},
				{Name: "Cyberethics - Professional Ethics and Codes of Conduct", Hours: 4},
				{Name: "Cyberethics - Ethics and Equity/Diversity", Hours: 4},
				{Name: "Cyberethics - Ethics and Law", Hours: 4},
				{Name: "Cyberethics - Ethics of Autonomy/Robots", Hours: 4},
				{Name: "Cyberethics - Ethics and Conflict", Hours: 4},
				{Name: "Cyberethics - Ethical Hacking", Hours: 4},
				{Name: "Cyberethics - Ethical Frameworks and Normative Theories", Hours: 4},
				{Name: "Cyber Policy - International Cyber Policy", Hours: 4},
				{Name: "Cyber Policy - U.S. Federal Cyber Policy", Hours: 4},
				{Name: "Cyber Policy - Global Impact", Hours: 4},
				{Name: "Cyber Policy - Cybersecurity and National Security Policy", Hours: 4},
				{Name: "Cyber Policy - National Economic Implications of Cybersecurity", Hours: 4},
				{Name: "Cyber Policy - New Adjacent Areas to Diplomacy", Hours: 4},
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

	// 4) Assign Cybersecurity Trainer / Awareness Specialist role to all existing degree programs
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
