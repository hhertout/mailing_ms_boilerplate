package mailer

import (
	"net/smtp"
	"os"
)

type Mailer struct {
	from      string
	smtpUrl   string
	plainAuth smtp.Auth
}

func NewMailer() *Mailer {
	m := &Mailer{
		from: os.Getenv("SMTP_FROM"),
	}
	m.smtpAuthentication()

	return m
}

func (m *Mailer) smtpAuthentication() {
	username := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	m.smtpUrl = smtpHost + ":" + smtpPort
	m.plainAuth = smtp.PlainAuth("", username, password, smtpHost)
}

func (m *Mailer) SendEmail(r *Request) error {
	r.SetHeaders()
	message := r.Headers + r.Body
	err := smtp.SendMail(m.smtpUrl, m.plainAuth, m.from, r.To, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
