package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"prosting/backend-gin/internal/models"
)

type ProjectRepository interface {
	Create(project models.Project) error
	FindAll() ([]models.Project, error)
	FindMainCategoriesByType(projectType string) ([]string, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(project models.Project) error {
	switch p := project.(type) {
	case models.ImageProject:
		return r.db.Create(&p).Error
	case models.MovieProject:
		return r.db.Create(&p).Error
	case models.MusicProject:
		return r.db.Create(&p).Error
	default:
		log.Fatal("Unknown project type")
		return nil
	}
}
func (r *projectRepository) FindAll() ([]models.Project, error) {
	var all []models.Project

	var images []models.ImageProject
	if err := r.db.Find(&images).Error; err != nil {
		return nil, err
	}
	for _, i := range images {
		all = append(all, i)
	}

	var videos []models.MovieProject
	if err := r.db.Find(&videos).Error; err != nil {
		return nil, err
	}
	for _, v := range videos {
		all = append(all, v)
	}

	var musics []models.MusicProject
	if err := r.db.Find(&musics).Error; err != nil {
		return nil, err
	}
	for _, m := range musics {
		all = append(all, m)
	}

	return all, nil
}

func (r *projectRepository) FindMainCategoriesByType(projectType string) ([]string, error) {
	var categories []string

	switch projectType {
	case "image":
		err := r.db.
			Model(&models.ImageProject{}).
			Distinct("main_category").
			Pluck("main_category", &categories).Error
		return categories, err
	case "music":
		err := r.db.
			Model(&models.MusicProject{}).
			Distinct("main_category").
			Pluck("main_category", &categories).Error
		return categories, err
	case "movie":
		err := r.db.
			Model(&models.MovieProject{}).
			Distinct("main_category").
			Pluck("main_category", &categories).Error
		return categories, err
	default:
		return nil, fmt.Errorf("unsupported project type: %s", projectType)
	}
}

func (r *projectRepository) FindByID(id uint) (models.Project, error) {
	var project models.Project
	if err := r.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (r *projectRepository) Update(project models.Project) error {
	switch p := project.(type) {
	case models.ImageProject:
		return r.db.Save(&p).Error
	case models.MovieProject:
		return r.db.Save(&p).Error
	case models.MusicProject:
		return r.db.Save(&p).Error
	default:
		log.Fatal("Unknown project type")
		return nil
	}
}

func (r *projectRepository) Delete(id uint) error {
	var project models.Project
	if err := r.db.First(&project, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&project).Error; err != nil {
		return err
	}
	return nil
}
