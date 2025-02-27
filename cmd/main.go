package main

import (
	"log"

	"github.com/joho/godotenv"

	"castracion-animal-api/internal/db"
	"castracion-animal-api/internal/handlers"
	"castracion-animal-api/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load environment variables: %s\n", err)
	}

	// Initialize the database
	if err := db.InitDB(); err != nil {
		log.Fatalf("Could not initialize the database: %s\n", err)
	}

	// Register routes
	userHandler := &handlers.UserController{}

	router := routes.RegisterRoutes(userHandler)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
