package main

import (
	"financial_tracker/app/config"
	"financial_tracker/app/infrastructure/db"
	"financial_tracker/app/infrastructure/http"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize the database
	database, err := db.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer database.Close()

	// Set up the router and start the server
	handler := http.NewHandler(database)
	router := handler.SetupRouter()
	router.Run(":8080")
}
