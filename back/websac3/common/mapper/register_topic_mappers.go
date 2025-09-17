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
