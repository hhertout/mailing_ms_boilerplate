package router

import (
	"mailer_ms/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", controllers.Test)
}
