package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"castracion-animal-api/internal/db"
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
	router := routes.RegisterRoutes()

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
