package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerTopicMappers() {
	RegisterMapFunc(
		func(knowledgeAreaModel *model.KnowledgeArea) (entity.KnowledgeArea, error) {
			var name string
			if len(knowledgeAreaModel.Names) > 0 {
				name = knowledgeAreaModel.Names[0].Name
			}
			return entity.KnowledgeArea{
				ID:   knowledgeAreaModel.ID,
				Name: name,
			}, nil
		},
	)

	// Mapper with language support for KnowledgeArea
	RegisterMapFunc(
		func(knowledgeAreaModel *model.KnowledgeArea) (entity.KnowledgeArea, error) {
			return mapKnowledgeAreaWithLanguage(knowledgeAreaModel, "en") // Default to English
		},
	)

	RegisterMapFunc(
		func(topicModel *model.Topic) (entity.Topic, error) {
			var err error
			var knowledgeArea entity.KnowledgeArea
			knowledgeArea, err = Map[model.KnowledgeArea, entity.KnowledgeArea](&topicModel.KnowledgeArea)
			if err != nil {
				return entity.Topic{}, err
			}

			var name string
			if len(topicModel.Names) > 0 {
				name = topicModel.Names[0].Name
			}

			return entity.Topic{
				ID:              topicModel.ID,
				Name:            name,
				KnowledgeAreaID: topicModel.KnowledgeAreaID,
				KnowledgeArea:   &knowledgeArea,
			}, nil
		},
	)

	// Mapper with language support - this will be overridden by specific language mappers
	RegisterMapFunc(
		func(topicModel *model.Topic) (entity.Topic, error) {
			return mapTopicWithLanguage(topicModel, "en") // Default to English
		},
	)

	RegisterMapFunc(
		func(topicEntity *entity.Topic) (response.ListTopicResponse, error) {
			var knowledgeAreaID uint
			var knowledgeAreaName string
			if topicEntity.KnowledgeArea != nil {
				knowledgeAreaID = topicEntity.KnowledgeArea.ID
				knowledgeAreaName = topicEntity.KnowledgeArea.Name
			}
			return response.ListTopicResponse{
				ID:                topicEntity.ID,
				Name:              topicEntity.Name,
				KnowledgeAreaID:   knowledgeAreaID,
				KnowledgeAreaName: knowledgeAreaName,
			}, nil
		},
	)
}

// Helper function to map topic with specific language
func mapTopicWithLanguage(topicModel *model.Topic, lang string) (entity.Topic, error) {
	var err error
	var knowledgeArea entity.KnowledgeArea
	knowledgeArea, err = Map[model.KnowledgeArea, entity.KnowledgeArea](&topicModel.KnowledgeArea)
	if err != nil {
		return entity.Topic{}, err
	}

	// Get name by language
	var name string
	for _, topicName := range topicModel.Names {
		if topicName.Lang == lang {
			name = topicName.Name
			break
		}
	}
	// If no name found for the language, use the first available
	if name == "" && len(topicModel.Names) > 0 {
		name = topicModel.Names[0].Name
	}

	return entity.Topic{
		ID:              topicModel.ID,
		Name:            name,
		KnowledgeAreaID: topicModel.KnowledgeAreaID,
		KnowledgeArea:   &knowledgeArea,
	}, nil
}

// Helper function to map knowledge area with specific language
func mapKnowledgeAreaWithLanguage(knowledgeAreaModel *model.KnowledgeArea, lang string) (entity.KnowledgeArea, error) {
	// Get name by language
	var name string
	for _, kaName := range knowledgeAreaModel.Names {
		if kaName.Lang == lang {
			name = kaName.Name
			break
		}
	}
	// If no name found for the language, use the first available
	if name == "" && len(knowledgeAreaModel.Names) > 0 {
		name = knowledgeAreaModel.Names[0].Name
	}

	return entity.KnowledgeArea{
		ID:   knowledgeAreaModel.ID,
		Name: name,
	}, nil
}

// Language-specific mapper functions
func GetTopicMapperWithLanguage(lang string) func(*model.Topic) (entity.Topic, error) {
	return func(topicModel *model.Topic) (entity.Topic, error) {
		return mapTopicWithLanguage(topicModel, lang)
	}
}

func GetKnowledgeAreaMapperWithLanguage(lang string) func(*model.KnowledgeArea) (entity.KnowledgeArea, error) {
	return func(knowledgeAreaModel *model.KnowledgeArea) (entity.KnowledgeArea, error) {
		return mapKnowledgeAreaWithLanguage(knowledgeAreaModel, lang)
	}
}
