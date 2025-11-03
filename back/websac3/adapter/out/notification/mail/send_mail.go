package mail

import (
	"time"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/logging"
	"websac3/common/mail"
)

type NotificationAdapter struct {
	emailSender          mail.ServerSMTP
	logger               logging.Logger
	createEmailPort      persistence.CreateEmailPort
	getEmailPort         persistence.GetEmailPort
	updateEmailPort      persistence.UpdateEmailPort
	persistenceManager   db.Manager
	isSenderWorking      bool
	pendingNotifications []entity.EmailNotification
}

func (n *NotificationAdapter) sendAsync() {
	var err error
	var chunkSize uint = 25
	var page uint = 0

	for {
		if err = n.persistenceManager.ExecuteNonTransactional(func(db db.Context) error {
			if n.pendingNotifications, err = n.getEmailPort.GetChunkNotSent(
				chunkSize,
				page,
				db,
			); err != nil {
				return err
			}

			return nil
		}); err != nil {
			n.logger.Error("Error al obtener notificaciones pendientes. Error: %s", err.Error())
			n.isSenderWorking = false
			n.pendingNotifications = nil
			return
		}

		n.isSenderWorking = true
		for _, emailToSend := range n.pendingNotifications {
			if err = n.emailSender.Send(&emailToSend); err != nil {
				n.logger.Error("Error al enviar correo electrónico. Error: %s", err.Error())
				break
			}

			var currentTime time.Time = time.Now()
			emailToSend.SentAt = &currentTime
			if err = n.persistenceManager.ExecuteInTransaction(func(tx db.Context) error {
				if err = n.updateEmailPort.Update(&emailToSend, tx); err != nil {
					return err
				}
				return nil
			}); err != nil {
				n.logger.Error("Error al actualizar estado del correo electrónico enviado. Error: %s", err.Error())
				break
			}

		}
		if err != nil {
			break
		}
	}
	n.isSenderWorking = false
}

func (n *NotificationAdapter) Send(notification *entity.EmailNotification, ctx db.Context) error {
	// Si tiene adjuntos, enviar inmediatamente (no guardar para envío asíncrono)
	if len(notification.Attachments) > 0 {
		// Enviar inmediatamente
		if err := n.emailSender.Send(notification); err != nil {
			n.logger.Error("Error al enviar correo electrónico con adjuntos. Error: %s", err.Error())
			return err
		}

		// Marcar como enviado y guardar en DB
		var currentTime time.Time = time.Now()
		notification.SentAt = &currentTime

		if err := n.createEmailPort.Create(notification, ctx); err != nil {
			n.logger.Error("Error al registrar notificacion de correo electrónico enviado. Error: %s", err.Error())
			return err
		}

		return nil
	}

	// Si no tiene adjuntos, usar el flujo normal (asíncrono)
	if err := n.createEmailPort.Create(notification, ctx); err != nil {
		n.logger.Error("Error al registrar notificacion de correo electrónico. Error: %s", err.Error())
		return err
	}

	if !n.isSenderWorking {
		go n.sendAsync()
	}

	return nil
}

func NewNotificationAdapter(
	emailSender mail.ServerSMTP,
	createEmailPort persistence.CreateEmailPort,
	updateEmailPort persistence.UpdateEmailPort,
	getEmailPort persistence.GetEmailPort,
	persistenceManager db.Manager,
	logger logging.Logger,
) *NotificationAdapter {
	na := &NotificationAdapter{
		emailSender:          emailSender,
		createEmailPort:      createEmailPort,
		updateEmailPort:      updateEmailPort,
		getEmailPort:         getEmailPort,
		persistenceManager:   persistenceManager,
		logger:               logger,
		pendingNotifications: []entity.EmailNotification{},
		isSenderWorking:      false,
	}
	go na.sendAsync()

	return na
}
