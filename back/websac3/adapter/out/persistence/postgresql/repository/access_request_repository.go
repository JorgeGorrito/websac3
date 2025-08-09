package repository

import (
	"errors"
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/constants"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/app/port/out/persistence/filter"
	"websac3/common/mapper"

	"gorm.io/gorm"
)

type AccessRequestRepository struct {
	Repository
	statusEnum enum.StatusEnum
}

func NewAccessRequestRepository(statusEnum enum.StatusEnum) *AccessRequestRepository {
	return &AccessRequestRepository{
		statusEnum: statusEnum,
	}
}

func (a *AccessRequestRepository) Create(accessRequestToSave *entity.AccessRequest, ctx _db.Context) error {
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

func (a *AccessRequestRepository) Update(accessRequestToUpdate *entity.AccessRequest, ctx _db.Context) error {
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

func (a *AccessRequestRepository) GetUnvalidatedEmailByToken(validationToken string, ctx _db.Context) (entity.AccessRequest, error) {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return entity.AccessRequest{}, err
	}

	var accessRequestFound entity.AccessRequest
	var accessRequest model.AccessRequest
	if err = dbCtx.DB().
		Model(&accessRequest).
		Joins("Applicant").
		Joins("Status").
		Joins("VerificationEmail").
		Where("validation_email_code = ?", validationToken).
		Where("is_verified = ?", false).
		First(&accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.AccessRequest{}, nil
		}
		return entity.AccessRequest{}, err
	}

	accessRequestFound, err = mapper.Map[model.AccessRequest, entity.AccessRequest](&accessRequest)
	if err != nil {
		return entity.AccessRequest{}, err
	}

	return accessRequestFound, nil
}

func (a *AccessRequestRepository) GetLastCreatedByIdentificationAndEmail(
	identificationTypeID uint,
	identificationNumber string,
	email string,
	ctx _db.Context,
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

func (a *AccessRequestRepository) GetAuthenticatedEmailByFilters(
	page, perPage uint,
	filters filter.Filters,
	ctx _db.Context,
) ([]entity.AccessRequest, int64, error) {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	pendingStatus, err := a.statusEnum.GetByName(constants.Pending)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.AccessRequest{}).
		Joins("Applicant").
		Joins("Applicant.HigherEducationInstitution").
		Joins("Applicant.HigherEducationInstitution.Municipality").
		Joins("Applicant.HigherEducationInstitution.Department").
		Joins("Status").
		Joins("VerificationEmail").
		Where("access_requests.is_verified = ?", true).
		Where(`"Status".id = ?`, pendingStatus.ID)

	dbCtx.DBSet(baseQuery)

	dbCtxFiltered, err := filters.Apply(dbCtx, psqlfilter.FiltersRegistry)
	if err != nil {
		return nil, 0, err
	}

	dbCtx, err = a.CastDbContext(dbCtxFiltered)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if err := dbCtx.DB().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Preload("Applicant.IdentificationType").
		Preload("Applicant.HigherEducationInstitution.Ownership").
		Preload("Applicant.HigherEducationInstitution.InstitutionalCategory").
		Preload("Applicant.HigherEducationInstitution.Department").
		Preload("Applicant.HigherEducationInstitution.Municipality").
		Preload("VerificationEmail").
		Preload("Status").
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var accessRequestsFound []model.AccessRequest
	if err := dbCtx.DB().Find(&accessRequestsFound).Error; err != nil {
		return nil, 0, err
	}

	var accessRequests []entity.AccessRequest
	for _, accessRequest := range accessRequestsFound {
		var accessRequestEntity entity.AccessRequest
		if accessRequestEntity, err = mapper.Map[model.AccessRequest, entity.AccessRequest](&accessRequest); err != nil {
			return nil, 0, err
		}
		accessRequests = append(accessRequests, accessRequestEntity)
	}

	return accessRequests, total, nil
}

func (a *AccessRequestRepository) GetByID(ID uint, ctx _db.Context) (entity.AccessRequest, error) {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return entity.AccessRequest{}, err
	}

	var accessRequest model.AccessRequest
	if err := dbCtx.DB().
		Model(&model.AccessRequest{}).
		Joins("Applicant").
		Joins("Applicant.HigherEducationInstitution").
		Joins("Applicant.HigherEducationInstitution.Municipality").
		Joins("Applicant.HigherEducationInstitution.Department").
		Joins("Status").
		Joins("VerificationEmail").
		Where("access_requests.id = ?", ID).
		First(&accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.AccessRequest{}, nil
		}
		return entity.AccessRequest{}, err
	}

	var accessRequestFound entity.AccessRequest
	if accessRequestFound, err = mapper.Map[model.AccessRequest, entity.AccessRequest](&accessRequest); err != nil {
		return entity.AccessRequest{}, err
	}

	return accessRequestFound, nil
}
