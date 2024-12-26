package routes

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

// RegisterRoutes registers all routes for the application
func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Define the root route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{Message: "Welcome to my Go API! 🚀"}
		json.NewEncoder(w).Encode(response)
	})

	// Add more routes here as needed

	return mux
}
