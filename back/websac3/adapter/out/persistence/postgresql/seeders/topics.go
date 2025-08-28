package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_TOPICS_SEED         = "adapter/out/persistence/postgresql/seeders/seeds/topics.json"
	DEFAULT_PATH_TOPICS_NAME_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/topic_names_en.json"
	DEFAULT_PATH_TOPICS_NAME_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/topic_names_es.json"
)

type topics struct{}

func Topics() Seeder {
	return &topics{}
}

func (s *topics) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var topicsToSeed []model.Topic = make([]model.Topic, 0)

	if err := decoder.Decode(DEFAULT_PATH_TOPICS_SEED, &topicsToSeed); err != nil {
		return err
	}

	for _, topic := range topicsToSeed {
		if err := dbCtx.DB().Create(&topic).Error; err != nil {
			return err
		}
	}

	var topicNamesToSeed []model.TopicName = make([]model.TopicName, 0)
	var topicNamesEsToSeed []model.TopicName = make([]model.TopicName, 0)
	var topicNamesEnToSeed []model.TopicName = make([]model.TopicName, 0)

	if err := decoder.Decode(DEFAULT_PATH_TOPICS_NAME_ES_SEED, &topicNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_TOPICS_NAME_EN_SEED, &topicNamesEnToSeed); err != nil {
		return err
	}

	topicNamesToSeed = append(topicNamesToSeed, topicNamesEsToSeed...)
	topicNamesToSeed = append(topicNamesToSeed, topicNamesEnToSeed...)

	for _, topicName := range topicNamesToSeed {
		if err := dbCtx.DB().Create(&topicName).Error; err != nil {
			return err
		}
	}

	return nil
}
