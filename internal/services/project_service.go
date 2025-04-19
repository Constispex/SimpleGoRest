package services

import (
	"fmt"
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/internal/repository"
)

type ProjectService interface {
	Create(project models.Project) error
	FindAll() ([]models.Project, error)
	FindMainCategoriesByType(projectType string) ([]string, error)
	FindByType(projectType string) ([]models.Project, error)
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

func (s *projectService) FindMainCategoriesByType(projectType string) ([]string, error) {
	categories, err := s.repo.FindMainCategoriesByType(projectType)
	if err != nil {
		return nil, fmt.Errorf("could not find categories: %w", err)
	}
	return categories, nil
}

func (s *projectService) FindByType(projectType string) ([]models.Project, error) {
	projects, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("could not find projects: %w", err)
	}

	var filteredProjects []models.Project
	for _, project := range projects {
		switch p := project.(type) {
		case models.ImageProject:
			if p.Type == models.StringToProjectType(projectType) {
				filteredProjects = append(filteredProjects, p)
			}
		case models.MovieProject:
			if p.Type == models.StringToProjectType(projectType) {
				filteredProjects = append(filteredProjects, p)
			}
		case models.MusicProject:
			if p.Type == models.StringToProjectType(projectType) {
				filteredProjects = append(filteredProjects, p)
			}
		default:
			return nil, fmt.Errorf("unknown project type: %T", project)
		}
	}

	return filteredProjects, nil
}
