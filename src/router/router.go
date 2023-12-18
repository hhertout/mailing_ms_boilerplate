package router

import (
	"mailer_ms/src/controllers"
	"mailer_ms/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	r := gin.Default()
	c := controllers.NewApiController()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ApiKeMiddleware())

	r.GET("/ping", c.Ping)
	r.POST("/api/mailer/helloworld", c.HelloWorldWithHtml)

	return r
}
