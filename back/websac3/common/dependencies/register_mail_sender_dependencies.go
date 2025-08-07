package dependencies

import (
	"os"
	amail "websac3/adapter/out/notification/mail"
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"
	"websac3/common/logging"
	"websac3/common/mail"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerMailSenderDependencies() {
	emailFrom := os.Getenv("MAIL_FROM")
	credentialsJsonPath := os.Getenv("MAIL_CREDENTIALS_PATH")
	tokenPath := os.Getenv("MAIL_TOKEN_PATH")

	m.binder.Bind(
		andi.GetAbstractType[notification.SendMailPort](),
		func() any {
			return amail.NewNotificationAdapter(
				mail.NewServerSMTP(
					emailFrom,
					credentialsJsonPath,
					tokenPath,
				),
				container.Inject[persistence.CreateEmailPort](),
				container.Inject[persistence.UpdateEmailPort](),
				container.Inject[persistence.GetEmailPort](),
				container.Inject[db.Manager](),
				container.Inject[logging.Logger](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.EmailPort](),
		func() any {
			return repository.NewEmailRepository(
				container.Inject[message.Provider](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateEmailPort](),
		func() any { return container.Inject[persistence.EmailPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateEmailPort](),
		func() any { return container.Inject[persistence.EmailPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetEmailPort](),
		func() any { return container.Inject[persistence.EmailPort]() },
	)

}
