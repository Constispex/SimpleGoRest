package services

import (
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/internal/repository"
)

// UserService kapselt die Geschäftslogik rund um die User.
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(u *models.User) (*models.User, error) {
	return s.repo.Save(u)
}
