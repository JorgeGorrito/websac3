package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
)

type cisoRole struct{}

func CISORole() Seeder { return &cisoRole{} }

func (s *cisoRole) Seed(ctx _db.Context) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}

	// 1) Create/find ProfessionalRole: CISO
	var role model.ProfessionalRole
	if err := dbCtx.DB().Where("name = ?", "CISO").First(&role).Error; err != nil {
		// Try to create if not exists
		role = model.ProfessionalRole{Name: "CISO"}
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

	// 2) Define KA expected for CISO (weights sum to 1.0)
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
			Name: "Organizational Security", Weight: 0.30,
			Topics: []topicCfg{
				{Name: "Security Governance and Policy - Organizational Context", Hours: 12},
				{Name: "Security Governance and Policy - Security Governance", Hours: 10},
				{Name: "Security Governance and Policy - Executive and Board-Level Communication", Hours: 8},
				{Name: "Security Governance and Policy - Management Policy", Hours: 8},
				{Name: "Security Program Management - Project Management", Hours: 8},
				{Name: "Security Program Management - Resource Management", Hours: 6},
				{Name: "Security Program Management - Security Metrics", Hours: 8},
				{Name: "Security Program Management - Assurance and Quality Control", Hours: 6},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.25,
			Topics: []topicCfg{
				{Name: "Risk Management - Risk Identification", Hours: 8},
				{Name: "Risk Management - Risk Assessment and Analysis", Hours: 10},
				{Name: "Risk Management - Insider Threats", Hours: 6},
				{Name: "Risk Management - Risk Measurement and Evaluation Models and Methodologies", Hours: 8},
				{Name: "Risk Management - Risk Control", Hours: 8},
				{Name: "Analytical Tools - Performance Measurement (Metrics)", Hours: 6},
				{Name: "Analytical Tools - Data Analytics", Hours: 6},
				{Name: "Analytical Tools - Security Intelligence", Hours: 6},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.20,
			Topics: []topicCfg{
				{Name: "Security Governance and Policy - Law, Ethics, and Compliance", Hours: 8},
				{Name: "Security Governance and Policy - Privacy", Hours: 6},
				{Name: "Cyber Law - Constitutional Foundations of Cyber Law", Hours: 6},
				{Name: "Cyber Law - Privacy Laws", Hours: 6},
				{Name: "Cyber Law - Data Security Law", Hours: 6},
				{Name: "Cyber Law - Anti-Hacking Laws", Hours: 4},
				{Name: "Cyber Law - Digital Evidence", Hours: 4},
				{Name: "Cyber Law - Digital Contracts", Hours: 4},
				{Name: "Cyber Law - Multinational Conventions (Agreements)", Hours: 4},
				{Name: "Cyber Law - Cross-Border Privacy and Data Security Laws", Hours: 6},
				{Name: "Cyberethics - Professional Ethics and Codes of Conduct", Hours: 4},
				{Name: "Cyberethics - Ethics and Law", Hours: 4},
			},
		},
		{
			Name: "Organizational Security", Weight: 0.15,
			Topics: []topicCfg{
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Incident Response", Hours: 8},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Disaster Recovery", Hours: 6},
				{Name: "Business Continuity, Disaster Recovery, and Incident Management - Business Continuity", Hours: 6},
				{Name: "Digital Forensics - Reporting, Incident Response and Management", Hours: 4},
				{Name: "Cybersecurity Planning - Strategic Planning", Hours: 6},
				{Name: "Cybersecurity Planning - Operational and Tactical Management", Hours: 4},
			},
		},
		{
			Name: "Data Security", Weight: 0.10,
			Topics: []topicCfg{
				{Name: "Data Privacy - Overview", Hours: 4},
				{Name: "Privacy - Defining Privacy", Hours: 3},
				{Name: "Privacy - Privacy Rights", Hours: 3},
				{Name: "Privacy - Safeguarding Privacy", Hours: 4},
				{Name: "Privacy - Privacy Norms and Attitudes", Hours: 3},
				{Name: "Privacy - Privacy Breaches", Hours: 3},
				{Name: "Privacy - Privacy in Societies", Hours: 3},
				{Name: "Social and Personal Privacy - Social Theories of Privacy", Hours: 3},
				{Name: "Social and Personal Privacy - Privacy and Security in Social Networks", Hours: 3},
				{Name: "Personal Data Privacy and Security - Sensitive Personal Data (SPD)", Hours: 4},
				{Name: "Personal Data Privacy and Security - Personal Tracking and Digital Footprint", Hours: 3},
				{Name: "Usable Security and Privacy - Privacy Policy", Hours: 3},
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

	// 4) Assign CISO role to all existing degree programs
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
