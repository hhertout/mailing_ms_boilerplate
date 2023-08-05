package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hhertout/go_mailing_ws.git/controllers"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", controllers.Test)
}
