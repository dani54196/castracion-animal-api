package repositories

import (
	"castracion-animal-api/internal/db"
	_ "castracion-animal-api/internal/db"
	"castracion-animal-api/internal/models"
	"errors"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	users []models.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []models.User{},
	}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	return r.users, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return models.User{}, errors.New("all fields are required")
	}
	user.ID = uint(len(r.users) + 1)
	r.users = append(r.users, user)
	return user, nil
}

func CreateUser(user *models.User) error {
	result := db.DB.Create(user)
	return result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
