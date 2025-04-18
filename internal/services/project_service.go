package services

import (
	"fmt"
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/internal/repository"
)

type ProjectService interface {
	Create(project models.Project) error
	FindAll() ([]models.Project, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) Create(p models.Project) error {
	fmt.Print("Service, Creating project: ", p)
	return s.repo.Create(p)
}

func (s *projectService) FindAll() ([]models.Project, error) {
	return s.repo.FindAll()
}
