package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/internal/services"
)

type ProjectHandler struct {
	svc services.ProjectService
}

func NewProjectHandler(service services.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: service}
}

func (h *ProjectHandler) UploadImageProject(c *gin.Context) {
	projectJson := c.PostForm("project")
	if projectJson == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing project data"})
		return
	}

	var project models.ImageProject
	if err := json.Unmarshal([]byte(projectJson), &project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project JSON"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	uploadDir := "./uploads/image"
	if err := c.SaveUploadedFile(files[0], uploadDir+"/"+files[0].Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	var filePaths []string
	for _, file := range files {
		if err := c.SaveUploadedFile(file, uploadDir+"/"+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		filePaths = append(filePaths, uploadDir+"/"+file.Filename)
	}

	project.ImageURLs = filePaths

	if err := h.svc.Create(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save project"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
}

func (h *ProjectHandler) Create(c *gin.Context) {
	// Begrenze die maximale Größe der Anfrage
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 20<<20) // 20 MB

	// Lese das JSON aus dem 'project'-Feld
	projectJSON := c.PostForm("project")
	if projectJSON == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'project' field"})
		return
	}

	// Unmarshale das JSON in eine generische Map, um den Typ zu bestimmen
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(projectJSON), &raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	projType, ok := raw["type"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'type' field"})
		return
	}

	// Verarbeite die Anfrage basierend auf dem Projekttyp
	switch projType {
	case "image":
		var img models.ImageProject
		if err := json.Unmarshal([]byte(projectJSON), &img); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image project"})
			return
		}

		// Extrahiere die hochgeladenen Dateien
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
			return
		}
		files := form.File["files"]

		// Verarbeite die Dateien (z. B. speichern)
		for _, file := range files {
			// Beispiel: Datei speichern
			c.SaveUploadedFile(file, "./uploads/"+file.Filename)
			// Optional: Pfad zur gespeicherten Datei im Projektobjekt speichern
			img.ImageURLs = append(img.ImageURLs, "/uploads/"+file.Filename)
		}

		if err := h.svc.Create(img); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image project"})
			return
		}
	// Weitere Fälle für 'movie' und 'music' können analog hinzugefügt werden
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown project type"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	projects, err := h.svc.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}
