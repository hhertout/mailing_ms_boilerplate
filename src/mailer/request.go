package mailer

import (
	"bytes"
	"fmt"
	"html/template"
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
	method  string
	To      []string
	Subject string
	Headers string
	Body    string
}

func NewRequest(subject string, to []string) *Request {
	if len(to) == 0 {
		to = []string{os.Getenv("SMTP_TO")}
	}

	r := &Request{
		To:      to,
		Subject: subject,
	}
	r.method = method.TEXT
	return r
}

func (r *Request) ParseHTMLTemplate(fileName string, data interface{}) (*Request, error) {
	r.method = method.HTML
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	t, err := template.ParseFiles(rootPath + "/templates/" + fileName)
	if err != nil {
		err := fmt.Errorf("invalid file path, giving : %s", rootPath+"/templates/"+fileName)
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return nil, err
	}
	r.Body = buf.String()

	return r, nil
}

func (r *Request) SetHeaders() *Request {
	headers := make(map[string]string)

	headers["From"] = os.Getenv("SMTP_FROM")
	headers["To"] = fmt.Sprint(r.To)
	headers["Subject"] = r.Subject
	headers["MIME-Version"] = "1.0"

	switch r.method {
	case method.HTML:
		headers["Content-Type"] = "text/html; charset=\"utf-8\""
	case method.TEXT:
		headers["Content-Type"] = "text/plain; charset=\"utf-8\""
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
