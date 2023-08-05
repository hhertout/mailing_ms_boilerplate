package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hhertout/go_mailing_ws.git/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	router.Routes(r)
	r.Run()
}
