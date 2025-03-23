package controllers

import (
	"log"
	"mailer_ms/internal/infra/repository"
	"mailer_ms/pkg/mailer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	mailer     *mailer.Mailer
	repository *repository.Repository
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
