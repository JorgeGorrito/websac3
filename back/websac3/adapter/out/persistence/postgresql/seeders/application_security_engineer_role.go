package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type applicationSecurityEngineerRole struct{}

func ApplicationSecurityEngineerRole() Seeder { return &applicationSecurityEngineerRole{} }

func (s *applicationSecurityEngineerRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: Application Security Engineer
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "Application Security Engineer").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "Application Security Engineer"}
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

	// 2) Define KA expected for Application Security Engineer (weights sum to 1.0)
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
			Name: "Software Security", Weight: 0.40,
			Topics: []topicCfg{
				{Name: "Design - Derivation of Security Requirements", Hours: 8},
				{Name: "Design - Specification of Security Requirements", Hours: 8},
				{Name: "Design - Software Development Life Cycle / Secure Development Life Cycle", Hours: 12},
				{Name: "Design - Programming Languages and Type-Safe Languages", Hours: 6},
				{Name: "Implementation - Input Validation and Representation Verification", Hours: 10},
				{Name: "Implementation - Correct Use of APIs", Hours: 8},
				{Name: "Implementation - Use of Security Features", Hours: 8},
				{Name: "Implementation - Verification of Time and State Relationships", Hours: 6},
				{Name: "Implementation - Proper Handling of Exceptions and Errors", Hours: 6},
				{Name: "Implementation - Robust Programming", Hours: 8},
				{Name: "Implementation - Encapsulation of Structures and Modules", Hours: 6},
				{Name: "Implementation - Consideration of the Environment", Hours: 6},
				{Name: "Analysis and Testing - Static and Dynamic Analysis", Hours: 10},
				{Name: "Analysis and Testing - Unit Testing", Hours: 8},
				{Name: "Analysis and Testing - Integration Testing", Hours: 8},
				{Name: "Analysis and Testing - Software Testing", Hours: 10},
				{Name: "Maintenance - Configuration", Hours: 6},
				{Name: "Maintenance - Patching and Vulnerability Life Cycle", Hours: 8},
				{Name: "Maintenance - Environment Verification", Hours: 6},
				{Name: "Maintenance - DevOps", Hours: 8},
				{Name: "Maintenance - Retirement/Decommissioning", Hours: 4},
				{Name: "Documentation - Security Documentation", Hours: 4},
				{Name: "Ethics - Ethical Issues in Software Development", Hours: 4},
				{Name: "Ethics - Social Aspects of Software Development", Hours: 4},
				{Name: "Legal Aspects - Legal Aspects of Software Development", Hours: 4},
				{Name: "Legal Aspects - Vulnerability Disclosure", Hours: 6},
			},
		},
		{
			Name: "Component Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "Component Design - Security in Component Design", Hours: 8},
				{Name: "Component Design - Secure Component Design Principles", Hours: 8},
				{Name: "Component Design - Component Identification", Hours: 6},
				{Name: "Component Design - Anti-Reverse Engineering Techniques", Hours: 6},
				{Name: "Component Design - Side Channel Attack Mitigation", Hours: 6},
				{Name: "Component Design - Anti-Tamper Technologies", Hours: 6},
				{Name: "Component Acquisition - Supply Chain Security", Hours: 8},
				{Name: "Component Testing - Unit Testing Principles", Hours: 6},
				{Name: "Component Testing - Security Testing", Hours: 8},
				{Name: "Component Reverse Engineering - Design Reverse Engineering", Hours: 6},
				{Name: "Component Reverse Engineering - Software Reverse Engineering", Hours: 8},
			},
		},
		{
			Name: "System Security", Weight: 0.20,
			Topics: []topicCfg{
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
				{Name: "System Control - Intrusion Detection", Hours: 4},
				{Name: "System Control - Attacks", Hours: 6},
				{Name: "System Control - Defenses", Hours: 6},
				{Name: "System Control - Auditing", Hours: 4},
				{Name: "System Control - Malware", Hours: 4},
				{Name: "System Control - Vulnerability Models", Hours: 6},
				{Name: "System Control - Penetration Testing", Hours: 6},
				{Name: "System Control - Forensics", Hours: 4},
				{Name: "System Control - Recovery and Resilience", Hours: 4},
				{Name: "System Testing - Requirements Validation", Hours: 4},
				{Name: "System Testing - Component Composition Validation", Hours: 4},
				{Name: "System Testing - Unit Testing versus System Testing", Hours: 4},
				{Name: "System Testing - Formal Verification of Systems", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "Information Storage Security - Database Security", Hours: 6},
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
		{
			Name: "Organizational Security", Weight: 0.05,
			Topics: []topicCfg{
				{Name: "Security Program Management - Project Management", Hours: 4},
				{Name: "Security Program Management - Resource Management", Hours: 4},
				{Name: "Security Program Management - Security Metrics", Hours: 4},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 4},
				{Name: "Analytical Tools - Performance Measurement (Metrics)", Hours: 4},
				{Name: "Analytical Tools - Data Analytics", Hours: 4},
				{Name: "Analytical Tools - Security Intelligence", Hours: 4},
				{Name: "Risk Management - Risk Identification", Hours: 4},
				{Name: "Risk Management - Risk Assessment and Analysis", Hours: 4},
				{Name: "Risk Management - Risk Control", Hours: 4},
				{Name: "Personnel Security - Third-Party Security", Hours: 4},
				{Name: "Personnel Security - Secure Hiring Practices", Hours: 4},
				{Name: "Personnel Security - Secure Termination Practices", Hours: 4},
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

	// 4) Assign Application Security Engineer role to all existing degree programs
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
