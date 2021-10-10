package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := newRouter()

	router.Logger.Fatal(router.Start(":8080"))
}
