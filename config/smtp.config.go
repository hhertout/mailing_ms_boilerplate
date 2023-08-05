package config

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
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

type Request struct {
  plainAuth smtp.Auth
  smtpUrl string
  headers string
  from string
  to []string
  subject string
  body string
}

func NewRequest(subject string, to []string) (*Request) {
  from := os.Getenv("SMTP_FROM")

  r := &Request {
    from: from, 
    to: to,
    subject: subject,
  }
  r.smtpAuthentification()

  return r
}

func (r *Request) smtpAuthentification() {
  username := os.Getenv("SMTP_USER")
  password := os.Getenv("SMTP_PASSWORD")
  smtpHost := os.Getenv("SMTP_HOST")
  smtpPort := os.Getenv("SMTP_PORT")

  r.smtpUrl = smtpHost + ":" + smtpPort
  r.plainAuth = smtp.PlainAuth("", username, password, smtpHost)
}

func (r *Request) ParseHTMLTemplate(fileName string, data interface{}) (error) {
  rootPath, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  t, err := template.ParseFiles(rootPath + "/templates/" + fileName)
  if err != nil {
    return err
  }

  buf := new(bytes.Buffer)
  if err = t.Execute(buf, data); err != nil {
    return err
  }
  r.body = buf.String()

  return nil
}

func (r *Request) SetHeaders() {
  headers := make(map[string]string)

  headers["From"] = r.from
  headers["To"] = fmt.Sprint(r.to)
  headers["Subject"] = r.subject
  headers["MIME-Version"] = "1.0"
  headers["Content-Type"] = "text/html; chartset=\"utf-8\""

  for k, v := range headers {
    r.headers += fmt.Sprintf("%s: %s \r\n", k, v)
  }
  r.headers += "\r\n"
}

func (r *Request) SendEmail() {
  r.SetHeaders()
  fmt.Printf("headers => %s, body => %s", r.headers, r.body)
  message := r.headers + r.body
  err := smtp.SendMail(r.smtpUrl, r.plainAuth, r.from, r.to, []byte(message))
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Email successfully sent to %s", r.to)
}



