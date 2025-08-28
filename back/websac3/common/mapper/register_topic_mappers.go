package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerTopicMappers() {
	RegisterMapFunc(
		func(knowledgeAreaModel *model.KnowledgeArea) (entity.KnowledgeArea, error) {
			return entity.KnowledgeArea{
				ID:   knowledgeAreaModel.ID,
				Name: knowledgeAreaModel.Names[0].Name,
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

			return entity.Topic{
				ID:              topicModel.ID,
				Name:            topicModel.Names[0].Name,
				KnowledgeAreaID: topicModel.KnowledgeAreaID,
				KnowledgeArea:   &knowledgeArea,
			}, nil
		},
	)

	RegisterMapFunc(
		func(topicEntity *entity.Topic) (response.ListTopicResponse, error) {
			return response.ListTopicResponse{
				ID:                topicEntity.ID,
				Name:              topicEntity.Name,
				KnowledgeAreaID:   topicEntity.KnowledgeArea.ID,
				KnowledgeAreaName: topicEntity.KnowledgeArea.Name,
			}, nil
		},
	)
}
