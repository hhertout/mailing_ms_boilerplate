package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mailer_ms/src/mailer"
	"net/http"
)

func (a *ApiController) HelloWorldWithHtml(c *gin.Context) {
	type data struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	var body MailRequest[data]

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	r := mailer.NewRequest(body.Subject, body.To)
	_, err := r.ParseHTMLTemplate("helloworld.html", body.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: Failed to parse html template with given variables",
		})
		return
	}

	a.mailer.SendEmail(r)

	c.JSON(http.StatusOK, gin.H{
		"message": "Email successfully sent",
	})
}
