package mapper

import (
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
			return role, nil
		},
	)
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
}
