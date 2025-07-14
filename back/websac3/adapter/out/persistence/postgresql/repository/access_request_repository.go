package repository

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"

	"gorm.io/gorm"
)

type AccessRequestRepository struct {
	Repository
}

func NewAccessRequestRepository(messageProvider message.Provider) *AccessRequestRepository {
	return &AccessRequestRepository{}
}

func (a *AccessRequestRepository) Create(accessRequestToSave *entity.AccessRequest, ctx persistence.Context) error {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var accessRequest model.AccessRequest
	if accessRequest, err = mapper.Map[entity.AccessRequest, model.AccessRequest](accessRequestToSave); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&accessRequest).
		Error; err != nil {
		return err
	}
	accessRequestToSave.ID = accessRequest.ID

	return nil
}

func (a *AccessRequestRepository) Update(accessRequestToUpdate *entity.AccessRequest, ctx persistence.Context) error {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var accessRequest model.AccessRequest
	if accessRequest, err = mapper.Map[entity.AccessRequest, model.AccessRequest](accessRequestToUpdate); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Model(&accessRequest).
		Where("id = ?", accessRequest.ID).
		Updates(accessRequest).
		Error; err != nil {
		return err
	}

	return nil
}

func (a *AccessRequestRepository) GetLastCreatedByIdentificationAndEmail(
	identificationTypeID uint,
	identificationNumber string,
	email string,
	ctx persistence.Context,
) (entity.AccessRequest, error) {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return entity.AccessRequest{}, err
	}

	var accessRequest model.AccessRequest
	var accessRequestFound entity.AccessRequest
	if err = dbCtx.DB().
		Model(&accessRequest).
		Joins("Applicant").
		Joins("Status").
		Joins("VerificationEmail").
		Where(`"Applicant".identification_type_id = ?`, identificationTypeID).
		Where(`"Applicant".identification_number = ?`, identificationNumber).
		Where(`"VerificationEmail".to_email = ?`, email).
		Order("access_requests.created_at desc").
		First(&accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.AccessRequest{}, nil
		}
		return entity.AccessRequest{}, err
	}

	if accessRequestFound, err = mapper.Map[model.AccessRequest, entity.AccessRequest](&accessRequest); err != nil {
		return entity.AccessRequest{}, err
	}

	return accessRequestFound, nil
}
