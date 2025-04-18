package repository

import (
	"gorm.io/gorm"
	"log"
	"prosting/backend-gin/internal/models"
)

type ProjectRepository interface {
	Create(project models.Project) error
	FindAll() ([]models.Project, error)
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
