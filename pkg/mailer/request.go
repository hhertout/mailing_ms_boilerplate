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

// Request represents an email request with necessary details.
type Request struct {
	method      string
	To          []string
	Subject     string
	Headers     string
	Attachement []string
	Body        string
}

// NewRequest creates a new email request with the given subject and recipients.
// If no recipients are provided, it defaults to the SMTP_TO environment variable.
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

// AddTemplateFromString sets the email body to the provided template string.
func (r *Request) AddTemplateFromString(tmpl string) *Request {
	r.Body = tmpl
	return r
}

// ParseHTMLTemplate parses an HTML template file and sets the email body.
// It takes the template file name and data to be injected into the template.
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

// SetHeaders sets the email headers based on the request details.
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

// AddAttachement adds a file attachment to the email request.
// It takes the file path of the attachment.
func (r *Request) AddAttachement(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		panic(err)
	}

	r.Attachement = append(r.Attachement, filePath)
	return nil
}

// SetPlainTextBody sets the email body to the provided plain text string.
func (r *Request) SetPlainTextBody(body string) *Request {
	r.method = method.TEXT
	r.Body = body
	return r
}
