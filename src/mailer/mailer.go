package mailer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
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

	if username == "" || password == "" || smtpHost == "" || smtpPort == "" {
		panic("SMTP credentials not set")
	}

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

func (m *Mailer) SendMailWithAttachment(r *Request) error {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	boundary := w.Boundary()

	headers := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=%s\r\n\r\n",
		m.from,
		strings.Join(r.To, ", "),
		r.Subject,
		boundary,
	)
	body.WriteString(headers)

	part, err := w.CreatePart(map[string][]string{
		"Content-Type": {"text/html; charset=utf-8"},
	})
	if err != nil {
		return fmt.Errorf("failed to create email body part: %w", err)
	}
	_, err = part.Write([]byte(r.Body))
	if err != nil {
		return fmt.Errorf("failed to write email body: %w", err)
	}

	for _, attachment := range r.Attachement {
		if err := m.parseAttachment(w, attachment); err != nil {
			return err
		}
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close MIME writer: %w", err)
	}

	err = smtp.SendMail(m.smtpUrl, m.plainAuth, m.from, r.To, body.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (m *Mailer) parseAttachment(w *multipart.Writer, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", file, err)
	}
	defer f.Close()

	// Lire le contenu du fichier
	fileData, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", file, err)
	}

	header := map[string][]string{
		"Content-Disposition":       {fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(file))},
		"Content-Type":              {"application/octet-stream"},
		"Content-Transfer-Encoding": {"base64"},
	}

	part, err := w.CreatePart(header)
	if err != nil {
		return fmt.Errorf("failed to create MIME part for file %s: %w", file, err)
	}

	encoder := base64.NewEncoder(base64.StdEncoding, part)
	defer encoder.Close()

	_, err = encoder.Write(fileData)
	if err != nil {
		return fmt.Errorf("failed to encode file %s: %w", file, err)
	}

	return nil
}
