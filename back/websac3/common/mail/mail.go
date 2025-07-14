package mail

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"mime"
	"os"
	"websac3/app/domain/entity"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type ServerSMTP struct {
	emailFrom string
	srv       *gmail.Service
}

func (s *ServerSMTP) Send(notification *entity.EmailNotification) error {
	to := notification.To
	subject := mime.BEncoding.Encode("UTF-8", notification.Subject)
	body := notification.Content

	message := []byte("From: " + s.emailFrom + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body)

	var msg gmail.Message
	msg.Raw = base64.URLEncoding.EncodeToString(message)

	_, err := s.srv.Users.Messages.Send("me", &msg).Do()
	return err
}
func NewServerSMTP(emailFrom string, credentialsJSONPath string, tokenPath string) *ServerSMTP {
	ctx := context.Background()

	credBytes, err := os.ReadFile(credentialsJSONPath)
	if err != nil {
		panic("No se pudo leer credentials.json: " + err.Error())
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

	return &ServerSMTP{
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
