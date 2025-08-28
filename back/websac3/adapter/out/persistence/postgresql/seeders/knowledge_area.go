package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_KNOWLEDGE_AREA_SEED = "adapter/out/persistence/postgresql/seeders/seeds/knowledge_area.json"

	DEFAULT_PATH_KNOWLEDGE_AREA_NAME_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/knowledge_area_names_es.json"
	DEFAULT_PATH_KNOWLEDGE_AREA_NAME_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/knowledge_area_names_en.json"
)

type knowledgeAreas struct{}

func KnowledgeAreas() Seeder {
	return &knowledgeAreas{}
}

func (s *knowledgeAreas) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var knowledgeAreasToSeed []model.KnowledgeArea = make([]model.KnowledgeArea, 0)

	if err := decoder.Decode(DEFAULT_PATH_KNOWLEDGE_AREA_SEED, &knowledgeAreasToSeed); err != nil {
		return err
	}

	for _, knowledgeArea := range knowledgeAreasToSeed {
		if err := dbCtx.DB().Create(&knowledgeArea).Error; err != nil {
			return err
		}
	}

	var knowledgeAreaNamesToSeed []model.KnowledgeAreaName = make([]model.KnowledgeAreaName, 0)
	var knowledgeAreaNamesEsToSeed []model.KnowledgeAreaName = make([]model.KnowledgeAreaName, 0)
	var knowledgeAreaNamesEnToSeed []model.KnowledgeAreaName = make([]model.KnowledgeAreaName, 0)

	if err := decoder.Decode(DEFAULT_PATH_KNOWLEDGE_AREA_NAME_ES_SEED, &knowledgeAreaNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_KNOWLEDGE_AREA_NAME_EN_SEED, &knowledgeAreaNamesEnToSeed); err != nil {
		return err
	}

	knowledgeAreaNamesToSeed = append(knowledgeAreaNamesToSeed, knowledgeAreaNamesEsToSeed...)
	knowledgeAreaNamesToSeed = append(knowledgeAreaNamesToSeed, knowledgeAreaNamesEnToSeed...)

	for _, knowledgeAreaName := range knowledgeAreaNamesToSeed {
		if err := dbCtx.DB().Create(&knowledgeAreaName).Error; err != nil {
			return err
		}
	}

	return nil
}
