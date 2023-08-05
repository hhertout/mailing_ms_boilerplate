package Controllers

import (
	GoMailer "mailer_ms/services"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	r := GoMailer.NewRequest("Hello World !", []string{"toto@gmail.com"})
	r.SetPlainTextBody("Send from Go Api").SendEmail()
}

func HelloWorldWithHtml(c *gin.Context) {
	templatePath := "helloworld.html"
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "toto",
		URL:  "http://go.dev",
	}
	r := GoMailer.NewRequest("Hello from go Controller", []string{"jane.doe@gmail.com"})
	r.ParseHTMLTemplate(templatePath, templateData).SendEmail()
}
