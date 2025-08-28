package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type TopicRepository struct {
	Repository
}

func NewTopicRepository() *TopicRepository {
	return &TopicRepository{}
}

func (r *TopicRepository) GetByNameAndLang(
	page uint,
	perPage uint,
	name string,
	lang string,
	ctx _db.Context,
) ([]entity.Topic, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Base con filtro por EXISTS
	base := dbCtx.DB().
		Model(&model.Topic{})

	sub := dbCtx.DB().Table("topic_names tn").
		Where("tn.topic_id = topics.id").
		Where("tn.lang = ? AND tn.name ILIKE ?", lang, "%"+name+"%")

	base = base.
		Where("EXISTS (?)", sub).
		Preload("Names", "lang = ? AND name ILIKE ?", lang, "%"+name+"%").
		Preload("KnowledgeArea.Names", "lang = ?", lang)

	dbCtx.DBSet(base)

	var count int64
	if err := dbCtx.DB().Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var topics []model.Topic
	if err := dbCtx.DB().Find(&topics).Error; err != nil {
		return nil, 0, err
	}

	var results []entity.Topic
	for _, t := range topics {
		topicEntity, err := mapper.Map[model.Topic, entity.Topic](&t)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, topicEntity)
	}

	return results, count, nil
}
