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
	from      string
	to        []string
	subject   string
	headers   string
	body      string
}

func NewRequest(subject string, to []string) *Request {
	r := &Request{
		from:    os.Getenv("SMTP_FROM"),
		to:      to,
		subject: subject,
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
	r.body = buf.String()

	return r
}

func (r *Request) SetHeaders() *Request {
	headers := make(map[string]string)

	headers["From"] = r.from
	headers["To"] = fmt.Sprint(r.to)
	headers["Subject"] = r.subject
	headers["MIME-Version"] = "1.0"

	switch r.method {
	case method.HTML:
		headers["Content-Type"] = "text/html; chartset=\"utf-8\""
	case method.TEXT:
		headers["Content-Type"] = "text/plain; chartset=\"utf-8\""
	}

	for k, v := range headers {
		r.headers += fmt.Sprintf("%s: %s \r\n", k, v)
	}
	r.headers += "\r\n"

	return r
}

func (r *Request) SetPlainTextBody(body string) *Request {
	r.method = method.TEXT
	r.body = body

	return r
}

func (r *Request) SendEmail() {
	r.SetHeaders()
	message := r.headers + r.body
	err := smtp.SendMail(r.smtpUrl, r.plainAuth, r.from, r.to, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Email successfully sent to %s", r.to)
}
