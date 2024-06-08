package main

import (
	"database/sql"
	"financial_tracker/app/infrastructure/db"
	"financial_tracker/app/infrastructure/http"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Prepare DB Config
	dbConfig := db.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// Initialize the database
	dbConn, err := sql.Open("postgres", dbConfig.ConnectionString())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	// Set up the router and start the server
	handler := http.NewHandler(dbConn)
	router := handler.SetupRouter()
	router.Run(":8080")
}
