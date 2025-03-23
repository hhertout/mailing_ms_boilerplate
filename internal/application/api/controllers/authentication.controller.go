package controllers

import (
	"mailer_ms/pkg/mailer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a ApiController) UpdatePasswordConfirmation(c *gin.Context) {
	type data struct {
		Email string `json:"email"`
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
	_, err := r.ParseHTMLTemplate("passwordUpdated.html", body.Body)
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

func (a ApiController) PasswordReinitialisation(c *gin.Context) {
	type data struct {
		Password string `json:"password"`
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
	_, err := r.ParseHTMLTemplate("passwordReinitialisation.html", body.Body)
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
