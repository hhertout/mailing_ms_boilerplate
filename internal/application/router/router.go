package router

import (
	"mailer_ms/internal/application/controllers"
	"mailer_ms/internal/application/middlewares"

	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	r := gin.Default()
	c := controllers.NewApiController()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ApiKeMiddleware())

	r.GET("/ping", c.Ping)

	r.GET("/api/mailer/list", c.GetMails)
	r.POST("/api/mailer/helloworld", c.HelloWorldWithHtml)

	r.POST("/api/mailer/password-updated", c.UpdatePasswordConfirmation)
	r.POST("/api/mailer/password-reset", c.PasswordReinitialisation)

	return r
}
