package services

import (
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/internal/repository"
)

// UserService kapselt die Geschäftslogik rund um die User.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService erstellt eine neue Instanz von UserService.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(u *models.User) (*models.User, error) {
	return s.repo.Save(u)
}
