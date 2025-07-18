package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerEmailMappers() {
	RegisterMapFunc(func(email *model.Email) (entity.EmailNotification, error) {
		return entity.EmailNotification{
			ID:        email.ID,
			To:        email.ToEmail,
			Subject:   email.Subject,
			Content:   email.Body,
			CreatedAt: email.CreatedAt,
			SentAt:    email.SentAt,
		}, nil
	})

	RegisterMapFunc(func(email *entity.EmailNotification) (model.Email, error) {
		return model.Email{
			ID:        email.ID,
			ToEmail:   email.To,
			Subject:   email.Subject,
			Body:      email.Content,
			CreatedAt: email.CreatedAt,
			SentAt:    email.SentAt,
		}, nil
	})

	RegisterMapFunc(func(email *request.ValidateEmailRequest) (command.ValidateEmailCommand, error) {
		return command.ValidateEmailCommand{
			ValidationToken: email.ValidationToken,
		}, nil
	})
}
