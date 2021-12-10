package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load ENVIRONMENT variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	// Build Container with Dig &  inject dependencies
	container := buildContainer()
	// Invoke the server with the container
	err := container.Invoke(server)

	if err != nil {
		log.Fatal(err)
	}
}
