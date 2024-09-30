package notification

import (
	"fmt"
	"net/smtp"

	"github.com/Rajendro1/Talenzen/config"
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
