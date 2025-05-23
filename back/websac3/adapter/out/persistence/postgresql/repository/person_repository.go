package repository

import (
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"
)

type PersonRepository struct{}

func (p *PersonRepository) Create(person *entity.Person, tx persistence.Transaction) error {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var personToSave models.Person
	if err := mapper.Map(person, &personToSave); err != nil {
		return err
	}
	fmt.Printf("person model: %+v\n", personToSave)

	if err := pgTx.Tx().Create(&personToSave).Error; err != nil {
		return err
	}
	person.ID = personToSave.ID

	return nil
}

func (p *PersonRepository) UpdateById(person *entity.Person, personID uint, tx persistence.Transaction) error {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var personToUpdate models.Person
	if err := mapper.Map(person, &personToUpdate); err != nil {
		return err
	}

	if err := pgTx.Tx().Where("id = ?", personID).Updates(personToUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (p *PersonRepository) GetByIdentificationNumber(identificationNumber string, tx persistence.Transaction) (entity.Person, error) {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return entity.Person{}, fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var person models.Person
	if err := pgTx.Tx().Where("identification_number = ?", identificationNumber).First(&person).Error; err != nil {
		return entity.Person{}, err
	}

	var personEntity entity.Person
	if err := mapper.Map(&person, &personEntity); err != nil {
		return entity.Person{}, err
	}

	return personEntity, nil
}
