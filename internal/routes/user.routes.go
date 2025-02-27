package routes

import (
	"castracion-animal-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *handlers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "all users"})
		})
		userRoutes.POST("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "user created"})
		})
	}
}
