package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type PersonRepository struct {
	Repository
}

func NewPersonRepository(messageProvider message.Provider) *PersonRepository {
	return &PersonRepository{}
}

func (p *PersonRepository) Create(personToSave *entity.Person, ctx _db.Context) error {
	dbCtx, err := p.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var person model.Person
	if person, err = mapper.Map[entity.Person, model.Person](personToSave); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&person).
		Error; err != nil {
		return err
	}
	personToSave.ID = person.ID

	return nil
}

func (p *PersonRepository) UpdateById(personToUpdate *entity.Person, personID uint, ctx _db.Context) error {
	dbCtx, err := p.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var person model.Person
	if person, err = mapper.Map[entity.Person, model.Person](personToUpdate); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Where("id = ?", personID).
		Updates(&person).Error; err != nil {
		return err
	}

	return nil
}

func (p *PersonRepository) GetByIdentificationNumber(identificationNumber string, ctx _db.Context) (entity.Person, error) {
	dbCtx, err := p.CastDbContext(ctx)
	if err != nil {
		return entity.Person{}, err
	}

	var person model.Person
	if err := dbCtx.DB().
		Model(&person).
		Where("identification_number = ?", identificationNumber).
		First(&person).Error; err != nil {
		return entity.Person{}, err
	}

	var personEntity entity.Person
	if personEntity, err = mapper.Map[model.Person, entity.Person](&person); err != nil {
		return entity.Person{}, err
	}

	return personEntity, nil
}
