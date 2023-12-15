package main

import (
	"fmt"
	"log"
	"mailer_ms/src/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.SetupRouter()
	fmt.Println("Starting server...")
	if err = r.Run(); err != nil {
		fmt.Println("Error on running server")
		fmt.Printf("Error: %s", err)

		return
	}
	fmt.Println("Server successfully started")
}
