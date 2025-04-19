package models

import "github.com/lib/pq"

type ProjectType string

const (
	ProjectTypeNone  ProjectType = "none"
	ProjectTypeMovie ProjectType = "movie"
	ProjectTypeAudio ProjectType = "music"
	ProjectTypeImage ProjectType = "image"
)

type Project interface {
	GetBase() ProjectBase
}

type ProjectBase struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	Title        string      `json:"title"`
	Year         int         `json:"year"`
	Description  string      `json:"description"`
	MainCategory string      `json:"mainCategory"`
	Type         ProjectType `json:"type"`
}

type MovieProject struct {
	ProjectBase
	VideoURL     *string        `json:"videoUrl,omitempty"`
	ThumbnailURL string         `json:"thumbnailUrl"`
	Tasks        pq.StringArray `gorm:"type:text[]" json:"tasks"`
}

func (p MovieProject) GetBase() ProjectBase {
	return p.ProjectBase
}

type MusicProject struct {
	ProjectBase
	AudioURL string `json:"audioUrl"`
}

func (p MusicProject) GetBase() ProjectBase {
	return p.ProjectBase
}

type ImageProject struct {
	ProjectBase
	ImageFolder string         `json:"imageFolder"`
	ImageURLs   pq.StringArray `gorm:"type:text[]" json:"imageUrls"`
	SubCategory string         `json:"subCategory"`
}

func (p ImageProject) GetBase() ProjectBase {
	return p.ProjectBase
}

func StringToProjectType(category string) ProjectType {
	switch category {
	case string(ProjectTypeMovie):
		return ProjectTypeMovie
	case string(ProjectTypeImage):
		return ProjectTypeImage
	case string(ProjectTypeAudio):
		return ProjectTypeAudio
	default:
		return ProjectTypeNone
	}
}
