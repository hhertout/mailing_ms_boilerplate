package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mailer_ms/src/mailer"
	"mailer_ms/src/repository"
	"net/http"
)

type ApiController struct {
	mailer     *mailer.Mailer
	repository *repository.Repository
}

type MailRequest[T interface{}] struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    T        `json:"body"`
}

func NewApiController() *ApiController {
	newRepository, err := repository.NewRepository(nil)
	if err != nil {
		log.Printf("Repositority initialisation failed : %s", err)
	}
	return &ApiController{
		mailer:     mailer.NewMailer(),
		repository: newRepository,
	}
}

func (a ApiController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
