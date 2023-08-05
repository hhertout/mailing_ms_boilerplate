package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hhertout/go_mailing_ws.git/router"
)

func main() {
	r := gin.Default()

  router.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
