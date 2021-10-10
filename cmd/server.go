package cmd

import (
	"github.com/joho/godotenv"
	"log"
)

func StartServer() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := newRouter()

	router.Logger.Fatal(router.Start(":8080"))
}
