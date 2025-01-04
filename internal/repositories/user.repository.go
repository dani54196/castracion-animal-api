package repositories

import (
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
