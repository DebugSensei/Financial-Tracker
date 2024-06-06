package main

import (
	"financial_tracker/infrastructure/db"
	"financial_tracker/infrastructure/http"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database
	db.Init()

	// Set up the router and start the server
	router := http.SetupRouter()
	router.Run(":8080")
}
