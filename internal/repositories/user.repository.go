package repositories

import (
	"castracion-animal-api/internal/db"
	"castracion-animal-api/internal/models"
)

func FindAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}
