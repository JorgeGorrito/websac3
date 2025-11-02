package mail

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"os"
	"time"
	"websac3/app/domain/entity"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type ServerSMTP interface {
	Send(notification *entity.EmailNotification) error
}

type serverSMTP struct {
	emailFrom string
	srv       *gmail.Service
}

func (s *serverSMTP) Send(notification *entity.EmailNotification) error {
	to := notification.To
	subject := mime.BEncoding.Encode("UTF-8", notification.Subject)
	body := notification.Content

	var message []byte
	var err error

	// Si no hay adjuntos, usar el formato simple
	if len(notification.Attachments) == 0 {
		message = []byte("From: " + s.emailFrom + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
			body)
	} else {
		// Si hay adjuntos, crear mensaje multipart
		message, err = s.createMultipartMessage(s.emailFrom, to, subject, body, notification.Attachments)
		if err != nil {
			return err
		}
	}

	var msg gmail.Message
	msg.Raw = base64.URLEncoding.EncodeToString(message)

	_, err = s.srv.Users.Messages.Send("me", &msg).Do()
	return err
}

func (s *serverSMTP) createMultipartMessage(from, to, subject, body string, attachments []entity.EmailAttachment) ([]byte, error) {
	var buf bytes.Buffer

	// Crear boundary manualmente
	boundary := "----=_NextPart_" + fmt.Sprintf("%d", time.Now().Unix())

	// Headers del mensaje principal
	buf.WriteString(fmt.Sprintf("From: %s\r\n", from))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", to))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n\r\n", boundary))

	// Parte del cuerpo HTML
	buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	buf.WriteString(body)
	buf.WriteString("\r\n\r\n")

	// Agregar adjuntos
	for _, attachment := range attachments {
		buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		buf.WriteString(fmt.Sprintf("Content-Type: %s\r\n", attachment.ContentType))
		buf.WriteString("Content-Transfer-Encoding: base64\r\n")
		buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", attachment.Filename))

		// Codificar datos en base64
		encodedData := base64.StdEncoding.EncodeToString(attachment.Data)

		// Dividir en líneas de 76 caracteres
		for i := 0; i < len(encodedData); i += 76 {
			end := i + 76
			if end > len(encodedData) {
				end = len(encodedData)
			}
			buf.WriteString(encodedData[i:end] + "\r\n")
		}
		buf.WriteString("\r\n")
	}

	// Cerrar el boundary
	buf.WriteString(fmt.Sprintf("--%s--\r\n", boundary))

	return buf.Bytes(), nil
}

var ErrSMTPNotConfigured = errors.New("SMTP server not configured")

type noOpServerSMTP struct{}

func (s *noOpServerSMTP) Send(notification *entity.EmailNotification) error {
	return ErrSMTPNotConfigured
}

func NewServerSMTP(emailFrom string, credentialsJSONPath string, tokenPath string) ServerSMTP {
	ctx := context.Background()

	credBytes, err := os.ReadFile(credentialsJSONPath)
	if err != nil {
		return &noOpServerSMTP{}
	}

	config, err := google.ConfigFromJSON(credBytes, gmail.GmailSendScope)
	if err != nil {
		panic("Error al parsear credentials.json: " + err.Error())
	}

	token, err := tokenFromFile(tokenPath)
	if err != nil {
		panic("No se pudo leer token.json: " + err.Error())
	}

	client := config.Client(ctx, token)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		panic("Error al crear servicio Gmail: " + err.Error())
	}

	return &serverSMTP{
		emailFrom: emailFrom,
		srv:       srv,
	}
}

func tokenFromFile(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	return &token, err
}
