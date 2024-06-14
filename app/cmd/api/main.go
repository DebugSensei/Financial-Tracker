package main

import (
	"context"
	"financial_tracker/app/config"
	"financial_tracker/app/infrastructure/db"
	"financial_tracker/app/ports"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Prepare Config
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
	handler := ports.NewHandler(database) // исправляем использование NewHandler
	router := handler.SetupRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Create a channel to listen for interrupt or terminate signals from the OS
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	log.Println("Server started on port 8080")

	// Wait for a signal to gracefully shut down the server
	<-stop
	log.Println("Shutting down server...")

	// Create a context with timeout to allow current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server gracefully stopped")
}
