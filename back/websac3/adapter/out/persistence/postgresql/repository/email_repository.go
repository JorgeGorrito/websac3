package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"
)

type EmailRepository struct {
	Repository
}

func NewEmailRepository(messageProvider message.Provider) *EmailRepository {
	return &EmailRepository{}
}

func (e *EmailRepository) Create(emailToCreate *entity.EmailNotification, ctx persistence.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var email model.Email
	if email, err = mapper.Map[entity.EmailNotification, model.Email](emailToCreate); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&email).
		Error; err != nil {
		return err
	}
	emailToCreate.ID = email.ID
	return nil
}

func (e *EmailRepository) GetChunkNotSent(chunkSize uint, page uint, ctx persistence.Context) ([]entity.EmailNotification, error) {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var emails []model.Email
	if err = dbCtx.DB().
		Model(&model.Email{}).
		Where("sent_at is null").
		Limit(int(chunkSize)).
		Offset(int(page * chunkSize)).
		Find(&emails).
		Error; err != nil {
		return nil, err
	}

	var emailNotifications []entity.EmailNotification
	for _, email := range emails {
		var emailNotification entity.EmailNotification
		if emailNotification, err = mapper.Map[model.Email, entity.EmailNotification](&email); err != nil {
			return nil, err
		}
		emailNotifications = append(emailNotifications, emailNotification)
	}

	return emailNotifications, nil

}

func (e *EmailRepository) Update(email *entity.EmailNotification, ctx persistence.Context) error {
	dbCtx, err := e.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var emailModel model.Email
	if emailModel, err = mapper.Map[entity.EmailNotification, model.Email](email); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Model(&emailModel).
		Where("id = ?", email.ID).
		Updates(emailModel).
		Error; err != nil {
		return err
	}

	return nil
}
