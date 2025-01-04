package services

import (
	"castracion-animal-api/internal/models"
	"castracion-animal-api/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}
