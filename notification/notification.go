package notification

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/smtp"

	"github.com/Rajendro1/Talenzen/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func SendEmail(to, subject, body string) error {
	from := config.SMTPUsername
	pass := config.SMTPPassword
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	auth := smtp.PlainAuth("", from, pass, config.SMTPServer)
	err := smtp.SendMail(config.SMTPServer+":"+config.SMTPPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("smtp error: %v", err)
	}
	return nil
}
func SendEmailOAuth(to, subject, body string) error {
	config := oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.GmailSendScope},
	}
	token := &oauth2.Token{AccessToken: "YOUR_ACCESS_TOKEN"}

	// Create a new HTTP client using the provided token.
	client := config.Client(context.Background(), token)
	gmailService, err := gmail.New(client)
	if err != nil {
		return fmt.Errorf("gmail service creation failed: %v", err)
	}

	var message gmail.Message
	emailBody := "From: 'me'\nTo: " + to + "\nSubject: " + subject + "\n\n" + body
	message.Raw = base64.URLEncoding.EncodeToString([]byte(emailBody))

	// Send the email.
	_, err = gmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return fmt.Errorf("gmail send error: %v", err)
	}
	return nil
}
