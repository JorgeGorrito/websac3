package repository

import (
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"
)

type AccessRequestRepository struct{}

func NewAccessRequestRepository() *AccessRequestRepository {
	return &AccessRequestRepository{}
}

func (a *AccessRequestRepository) Create(accessRequest *entity.AccessRequest, tx persistence.Transaction) error {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var accessRequestToSave models.AccessRequest
	if err := mapper.Map(accessRequest, &accessRequestToSave); err != nil {
		return err
	}
	fmt.Printf("create access request entity: %+v\n", accessRequestToSave)

	if err := pgTx.Tx().Create(&accessRequestToSave).Error; err != nil {
		return err
	}
	accessRequest.ID = accessRequestToSave.ID

	return nil
}

func (a *AccessRequestRepository) GetLastCreatedPersonIdentificationNumber(identificationNumber string, tx persistence.Transaction) (entity.AccessRequest, error) {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return entity.AccessRequest{}, fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var accessRequest models.AccessRequest
	var accessRequestFound entity.AccessRequest
	if err := pgTx.Tx().Table("access_requests a").
		Joins("inner join people p on a.applicant_id = p.id").
		Where("p.identification_number = ?", identificationNumber).
		Order("access_requests.created_at desc").Error; err != nil {
		return entity.AccessRequest{}, err
	}

	if err := mapper.Map(&accessRequest, &accessRequestFound); err != nil {
		return entity.AccessRequest{}, err
	}

	return accessRequestFound, nil
}
