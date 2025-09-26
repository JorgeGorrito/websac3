package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type devsecopsRole struct{}

func DevSecOpsRole() Seeder { return &devsecopsRole{} }

func (s *devsecopsRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: DevSecOps
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "DevSecOps").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "DevSecOps"}
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

	// 2) Define KA expected for DevSecOps (weights sum to 1.0)
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
			Name: "Software Security", Weight: 0.35,
			Topics: []topicCfg{
				{Name: "Design - Software Development Life Cycle / Secure Development Life Cycle", Hours: 12},
				{Name: "Implementation - Input Validation and Representation Verification", Hours: 10},
				{Name: "Analysis and Testing - Static and Dynamic Analysis", Hours: 10},
				{Name: "Maintenance - DevOps", Hours: 8},
			},
		},
		{
			Name: "System Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "System Control - Intrusion Detection", Hours: 8},
				{Name: "System Control - Vulnerability Models", Hours: 8},
				{Name: "System Control - Recovery and Resilience", Hours: 8},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Security Program Management - Security Metrics", Hours: 6},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 8},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 8},
			},
		},
		{
			Name: "Data Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Data Integrity and Authentication - Password Storage Techniques", Hours: 6},
				{Name: "Information Storage Security - Database Security", Hours: 8},
				{Name: "Data Privacy - Overview", Hours: 6},
			},
		},
		{
			Name: "Connection Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "Secure Communication Protocols - TLS Attacks", Hours: 6},
				{Name: "Network Defense - Firewalls and Virtual Private Networks (VPNs)", Hours: 8},
				{Name: "Network Defense - Network Monitoring", Hours: 6},
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

	// 4) Assign DevSecOps role to all existing degree programs
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
