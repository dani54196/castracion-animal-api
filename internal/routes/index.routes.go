package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

// RegisterRoutes registers all routes for the application
func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// CORS
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	// Define the root route
	router.GET("/", func(c *gin.Context) {
		response := Response{Message: "Connected to the API"}
		c.JSON(200, response)
	})

	return router
}
