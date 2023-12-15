package controllers

import (
	"github.com/gin-gonic/gin"
	"mailer_ms/src/mailer"
	"net/http"
)

func (a *ApiController) HelloWorldWithHtml(c *gin.Context) {
	templatePath := "helloworld.html"
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "toto",
		URL:  "http://go.dev",
	}

	r := mailer.NewRequest("Hello from go Controller", []string{"jane.doe@gmail.com"})
	r.ParseHTMLTemplate(templatePath, templateData)

	a.mailer.SendEmail(r)

	c.JSON(http.StatusOK, gin.H{
		"message": "Email successfully sent",
	})
}
