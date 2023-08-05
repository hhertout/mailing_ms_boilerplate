package config

import (
	"fmt"
	"net/smtp"
	"os"
)

func SmtpAuthentification() (smtp.Auth, string) {
  username := os.Getenv("SMTP_USER")
  password := os.Getenv("SMTP_PASSWORD")
  smtpHost := os.Getenv("SMTP_HOST")
  smtpPort := os.Getenv("SMTP_PORT")
  smtpUrl := smtpHost + ":" + smtpPort
  auth := smtp.PlainAuth("", username, password, smtpHost)
  return auth, smtpUrl
}


func SetHeaders(from *string, to *[]string, subject *string) (map[string]string) {
  headers := make(map[string]string)

  headers["From"] = *from
  headers["To"] = fmt.Sprint(to)
  headers["Subject"] = *subject
  headers["MIME-Version"] = "1.0"
  headers["Content-Type"] = "text/plain; chartset=\"utf-8\""
  
  return headers
}



