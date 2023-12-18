package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mailer_ms/src/mailer"
	"net/http"
)

func (a ApiController) HelloWorldWithHtml(c *gin.Context) {
	type data struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	var body MailRequest[data]

	if err := c.BindJSON(&body); err != nil {
		_ = a.repository.SaveWithError(body.To, body.Subject, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	r := mailer.NewRequest(body.Subject, body.To)
	_, err := r.ParseHTMLTemplate("helloworld.html", body.Body)
	if err != nil {
		_ = a.repository.SaveWithError(body.To, body.Subject, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: Failed to parse html template with given variables",
		})
		return
	}

	if err = a.mailer.SendEmail(r); err != nil {
		_ = a.repository.SaveWithError(body.To, body.Subject, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server Error",
		})
	}

	_ = a.repository.SaveWithoutError(body.To, body.Subject)
	c.JSON(http.StatusOK, gin.H{
		"message": "Email successfully sent",
	})
}

func (a ApiController) GetMails(c *gin.Context) {
	res, err := a.repository.ListMail()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mails": res,
	})
}
