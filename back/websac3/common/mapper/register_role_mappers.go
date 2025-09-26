package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerProfessionalRoleMappers() {
	RegisterMapFunc(
		func(roleModel *model.ProfessionalRole) (entity.ProfessionalRole, error) {
			role := entity.ProfessionalRole{
				ID:   roleModel.ID,
				Name: roleModel.Name,
			}

			var kaExpected []entity.KnowledgeAreaExpected
			for _, ka := range roleModel.KnowledgeAreas {
				// Map KnowledgeArea
				var eKA entity.KnowledgeArea
				if len(ka.KnowledgeArea.Names) > 0 {
					eKA = entity.KnowledgeArea{
						ID:   ka.KnowledgeArea.ID,
						Name: ka.KnowledgeArea.Names[0].Name,
					}
				}

				var topicsExpected []entity.TopicExpected
				for _, t := range ka.Topics {
					var eTopic entity.Topic
					if len(t.Topic.Names) > 0 {
						eTopic = entity.Topic{
							ID:              t.Topic.ID,
							Name:            t.Topic.Names[0].Name,
							KnowledgeAreaID: t.Topic.KnowledgeAreaID,
						}
					}
					topicsExpected = append(topicsExpected, entity.TopicExpected{
						ID:         t.ID,
						TopicID:    t.TopicID,
						Topic:      &eTopic,
						LearnHours: float32(t.StudyHours),
					})
				}

				kaExpected = append(kaExpected, entity.KnowledgeAreaExpected{
					ID:              ka.ID,
					KnowledgeAreaID: ka.KnowledgeAreaID,
					KnowledgeArea:   &eKA,
					TopicExpected:   topicsExpected,
					PriorityWeight:  ka.PriorityWeight,
				})
			}

			role.KnowledgeAreaExpected = kaExpected

			var degreePrograms []entity.DegreeProgram
			for _, dpm := range roleModel.DegreePrograms {
				dpe, err := Map[model.DegreeProgram, entity.DegreeProgram](&dpm)
				if err != nil {
					return entity.ProfessionalRole{}, err
				}
				degreePrograms = append(degreePrograms, dpe)
			}
			role.DegreePrograms = degreePrograms

			return role, nil
		},
	)

	// Entity to Response mapper for ProfessionalRole
	RegisterMapFunc(func(role *entity.ProfessionalRole) (response.ListProfessionalRoleResponse, error) {
		return response.ListProfessionalRoleResponse{
			ID:   role.ID,
			Name: role.Name,
		}, nil
	})
}

func registerRoleMappers() {
	RegisterMapFunc(func(role *model.Role) (entity.Role, error) {
		var err error

		var permissions []entity.Permission
		for _, permission := range role.Permissions {
			var permissionMapped entity.Permission
			permissionMapped, err = Map[model.Permission, entity.Permission](&permission)
			if err != nil {
				return entity.Role{}, err
			}
			permissions = append(permissions, permissionMapped)
		}

		return entity.Role{
			ID:          role.ID,
			Name:        role.Name,
			Permissions: permissions,
		}, nil
	})

	RegisterMapFunc(func(role *entity.Role) (model.Role, error) {
		return model.Role{
			ID:   role.ID,
			Name: role.Name,
		}, nil
	})

	// Entity to Response mapper
	RegisterMapFunc(func(role *entity.Role) (response.ListRoleResponse, error) {
		return response.ListRoleResponse{
			ID:   role.ID,
			Name: role.Name,
		}, nil
	})
}

// GetProfessionalRoleMapperWithLanguage returns a language-specific mapper for ProfessionalRole
func GetProfessionalRoleMapperWithLanguage(lang string) func(*model.ProfessionalRole) (entity.ProfessionalRole, error) {
	return func(roleModel *model.ProfessionalRole) (entity.ProfessionalRole, error) {
		role := entity.ProfessionalRole{
			ID:   roleModel.ID,
			Name: roleModel.Name,
		}

		var kaExpected []entity.KnowledgeAreaExpected
		for _, ka := range roleModel.KnowledgeAreas {
			// Map KnowledgeArea with language support
			var eKA entity.KnowledgeArea
			if len(ka.KnowledgeArea.Names) > 0 {
				// Find name by language, fallback to first available
				var name string
				for _, nameObj := range ka.KnowledgeArea.Names {
					if nameObj.Lang == lang {
						name = nameObj.Name
						break
					}
				}
				if name == "" && len(ka.KnowledgeArea.Names) > 0 {
					name = ka.KnowledgeArea.Names[0].Name
				}

				eKA = entity.KnowledgeArea{
					ID:   ka.KnowledgeArea.ID,
					Name: name,
				}
			}

			var topicsExpected []entity.TopicExpected
			for _, t := range ka.Topics {
				var eTopic entity.Topic
				if len(t.Topic.Names) > 0 {
					// Find name by language, fallback to first available
					var name string
					for _, nameObj := range t.Topic.Names {
						if nameObj.Lang == lang {
							name = nameObj.Name
							break
						}
					}
					if name == "" && len(t.Topic.Names) > 0 {
						name = t.Topic.Names[0].Name
					}

					eTopic = entity.Topic{
						ID:              t.Topic.ID,
						Name:            name,
						KnowledgeAreaID: t.Topic.KnowledgeAreaID,
					}
				}
				topicsExpected = append(topicsExpected, entity.TopicExpected{
					ID:         t.ID,
					TopicID:    t.TopicID,
					Topic:      &eTopic,
					LearnHours: float32(t.StudyHours),
				})
			}

			kaExpected = append(kaExpected, entity.KnowledgeAreaExpected{
				ID:              ka.ID,
				KnowledgeAreaID: ka.KnowledgeAreaID,
				KnowledgeArea:   &eKA,
				TopicExpected:   topicsExpected,
				PriorityWeight:  ka.PriorityWeight,
			})
		}

		role.KnowledgeAreaExpected = kaExpected

		var degreePrograms []entity.DegreeProgram
		for _, dpm := range roleModel.DegreePrograms {
			dpe, err := Map[model.DegreeProgram, entity.DegreeProgram](&dpm)
			if err != nil {
				return entity.ProfessionalRole{}, err
			}
			degreePrograms = append(degreePrograms, dpe)
		}
		role.DegreePrograms = degreePrograms

		return role, nil
	}
}
