package controllers

import (
	"github.com/gin-gonic/gin"
	"mailer_ms/src/mailer"
	"net/http"
)

type ApiController struct {
	mailer *mailer.Mailer
}

type MailRequest[T interface{}] struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    T        `json:"body"`
}

func NewApiController() *ApiController {
	return &ApiController{mailer: mailer.NewMailer()}
}

func (a *ApiController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
