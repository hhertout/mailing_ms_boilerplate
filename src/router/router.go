package router

import (
	"mailer_ms/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	c := controllers.NewApiController()

	r.GET("/ping", c.Ping)
	r.POST("/api/mailer/helloworld", c.HelloWorldWithHtml)

	return r
}
