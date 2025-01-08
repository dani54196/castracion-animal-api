package routes

import (
	"castracion-animal-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *handlers.UserController) {
	userRoutes := handlers.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
	}
}
