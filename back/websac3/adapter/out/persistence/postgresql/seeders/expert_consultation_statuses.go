package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_EXPERT_CONSULTATION_STATUS_SEED          = "adapter/out/persistence/postgresql/seeders/seeds/expert_consultation_statuses.json"
	DEFAULT_PATH_EXPERT_CONSULTATION_STATUS_NAMES_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/expert_consultation_status_names_en.json"
	DEFAULT_PATH_EXPERT_CONSULTATION_STATUS_NAMES_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/expert_consultation_status_names_es.json"
)

type expertConsultationStatuses struct{}

func ExpertConsultationStatuses() Seeder {
	return &expertConsultationStatuses{}
}

func (e *expertConsultationStatuses) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var statusesToSeed []model.ExpertConsultationStatus = make([]model.ExpertConsultationStatus, 0)

	if err := decoder.Decode(DEFAULT_PATH_EXPERT_CONSULTATION_STATUS_SEED, &statusesToSeed); err != nil {
		return err
	}

	for _, status := range statusesToSeed {
		if err := dbCtx.DB().Create(&status).Error; err != nil {
			return err
		}
	}

	var statusNamesToSeed []model.ExpertConsultationStatusName = make([]model.ExpertConsultationStatusName, 0)
	var statusNamesEsToSeed []model.ExpertConsultationStatusName = make([]model.ExpertConsultationStatusName, 0)
	var statusNamesEnToSeed []model.ExpertConsultationStatusName = make([]model.ExpertConsultationStatusName, 0)

	if err := decoder.Decode(DEFAULT_PATH_EXPERT_CONSULTATION_STATUS_NAMES_ES_SEED, &statusNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_EXPERT_CONSULTATION_STATUS_NAMES_EN_SEED, &statusNamesEnToSeed); err != nil {
		return err
	}

	statusNamesToSeed = append(statusNamesToSeed, statusNamesEsToSeed...)
	statusNamesToSeed = append(statusNamesToSeed, statusNamesEnToSeed...)

	for _, statusName := range statusNamesToSeed {
		if err := dbCtx.DB().Create(&statusName).Error; err != nil {
			return err
		}
	}

	return nil
}
