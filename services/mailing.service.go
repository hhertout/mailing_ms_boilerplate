package GoMailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

var method = struct {
	TEXT string
	HTML string
}{
	TEXT: "text",
	HTML: "html",
}

type Request struct {
	method    string
	plainAuth smtp.Auth
	smtpUrl   string
	From      string
	To        []string
	Subject   string
	Headers   string
	Body      string
}

func NewRequest(subject string, to []string) *Request {
	r := &Request{
		From:    os.Getenv("SMTP_FROM"),
		To:      to,
		Subject: subject,
	}
	r.method = method.TEXT
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

func (r *Request) ParseHTMLTemplate(fileName string, data interface{}) *Request {
	r.method = method.HTML
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles(rootPath + "/templates/" + fileName)
	if err != nil {
		fmt.Printf("Invalid file path, giving : %s", rootPath+"/templates/"+fileName)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}
	r.Body = buf.String()

	return r
}

func (r *Request) SetHeaders() *Request {
	headers := make(map[string]string)

	headers["From"] = r.From
	headers["To"] = fmt.Sprint(r.To)
	headers["Subject"] = r.Subject
	headers["MIME-Version"] = "1.0"

	switch r.method {
	case method.HTML:
		headers["Content-Type"] = "text/html; chartset=\"utf-8\""
	case method.TEXT:
		headers["Content-Type"] = "text/plain; chartset=\"utf-8\""
	}

	for k, v := range headers {
		r.Headers += fmt.Sprintf("%s: %s \r\n", k, v)
	}
	r.Headers += "\r\n"

	return r
}

func (r *Request) SetPlainTextBody(body string) *Request {
	r.method = method.TEXT
	r.Body = body

	return r
}

func (r *Request) SendEmail() {
	r.SetHeaders()
	message := r.Headers + r.Body
	err := smtp.SendMail(r.smtpUrl, r.plainAuth, r.From, r.To, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Email successfully sent to %s", r.To)
}
