package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/internal/services"
)

type ProjectHandler struct {
	svc services.ProjectService
}

func NewProjectHandler(service services.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: service}
}

func (h *ProjectHandler) Create(c *gin.Context) {
	// Begrenze die maximale Größe der Anfrage
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 20<<20) // 20 MB

	projectBytes, err := readProjectFile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read project file"})
		return
	}

	projectType, err := parseProjectType(projectBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project type"})
		return
	}

	// Lese den Inhalt des 'project'-Feldes
	switch projectType {
	case "image":
		if err := h.handleImageProject(c, projectBytes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	case "movie":
		if err := h.handleMovieProject(c, projectBytes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	case "music":
		if err := h.handleMusicProject(c, projectBytes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	// Weitere Fälle für 'movie' und 'music' können analog hinzugefügt werden
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown project type"})
		return
	}

	c.Status(http.StatusCreated)
}

func parseProjectType(projectBytes []byte) (string, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(projectBytes, &raw); err != nil {
		return "", err
	}
	projType, ok := raw["type"].(string)
	if !ok {
		return "", fmt.Errorf("missing or invalid 'type' field")
	}
	return projType, nil

}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	projectType := c.Query("type")
	if projectType != "" {
		projects, err := h.svc.FindByType(projectType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, projects)
		return
	}
	projects, err := h.svc.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetCategories(c *gin.Context) {
	projectType := c.Query("type")

	categories, err := h.svc.FindMainCategoriesByType(projectType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *ProjectHandler) handleImageProject(c *gin.Context, data []byte) error {
	var img models.ImageProject
	if err := json.Unmarshal(data, &img); err != nil {
		return fmt.Errorf("invalid image project data: %w", err)
	}

	categoryFolder := filepath.Join("uploads", img.MainCategory, img.SubCategory)

	savedFiles, err := saveToDisk(c, categoryFolder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image project: " + err.Error()})
		return err
	}
	img.ImageURLs = append(img.ImageURLs, savedFiles...)
	if err := h.svc.Create(img); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image project"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image project created successfully"})
	return nil
}

func (h *ProjectHandler) handleMovieProject(c *gin.Context, data []byte) error {
	var movie models.MovieProject
	if err := json.Unmarshal(data, &movie); err != nil {
		return fmt.Errorf("invalid movie project data: %w", err)
	}

	if movie.VideoURL == nil || *movie.VideoURL == "" {

		categoryFolder := filepath.Join("uploads", movie.MainCategory)

		savedFiles, err := saveToDisk(c, categoryFolder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save movie project: " + err.Error()})
			return err
		}
		movie.VideoURL = &savedFiles[0]
	} else {
		// Wenn die URL bereits gesetzt ist, speichern wir sie nicht erneut
		fmt.Println("Video URL already set, not saving to disk.")
	}
	if err := h.svc.Create(movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save movie project"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie project created successfully"})
	return nil
}

func (h *ProjectHandler) handleMusicProject(c *gin.Context, data []byte) error {
	var music models.MusicProject
	if err := json.Unmarshal(data, &music); err != nil {
		return fmt.Errorf("invalid music project data: %w", err)
	}

	categoryFolder := filepath.Join("uploads", music.MainCategory)

	savedFiles, err := saveToDisk(c, categoryFolder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save music project: " + err.Error()})
		return err
	}
	music.AudioURL = savedFiles[0]
	if err := h.svc.Create(music); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save music project"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Music project created successfully"})
	return nil
}

func saveToDisk(c *gin.Context, categoryFolder string) ([]string, error) {
	form := c.Request.MultipartForm
	files := form.File["files"]
	var urls []string

	for _, file := range files {
		// Beispiel: Datei speichern
		err := c.SaveUploadedFile(file, filepath.Join(categoryFolder, file.Filename))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file: " + err.Error()})
			return nil, err
		}
		urls = append(urls, filepath.Join(categoryFolder, file.Filename))
	}

	fmt.Println("Saved files:", urls)
	return urls, nil
}

func readProjectFile(c *gin.Context) ([]byte, error) {
	// Lese den Inhalt des 'project'-Feldes
	fileHeader, err := c.FormFile("project")
	if err != nil {
		return nil, err
	}

	projectFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(projectFile multipart.File) {
		err := projectFile.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}(projectFile)

	projectBytes, err := io.ReadAll(projectFile)
	if err != nil {
		return nil, err
	}

	return projectBytes, nil
}
