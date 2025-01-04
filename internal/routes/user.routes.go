package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := handlers.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
	}
}
