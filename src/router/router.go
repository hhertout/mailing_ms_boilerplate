package router

import (
	"mailer_ms/src/config"
	"mailer_ms/src/controllers"

	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	r := gin.Default()
	c := controllers.NewApiController()

	r.Use(config.CORSMiddleware())

	r.GET("/ping", c.Ping)
	r.POST("/api/mailer/helloworld", c.HelloWorldWithHtml)

	return r
}
