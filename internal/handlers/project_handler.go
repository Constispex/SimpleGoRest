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

func (h *ProjectHandler) Create(c *gin.Context) {
	// 1. JSON komplett als []byte lesen
	var raw map[string]interface{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// 2. Projekttyp prüfen
	projType, ok := raw["type"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'type' field"})
		return
	}

	// 3. Den JSON-Body nochmal serialisieren (um ihn in struct zu parsen)
	jsonBody, err := json.Marshal(raw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot re-marshal body"})
		return
	}

	// 4. Switch auf Typ → spezifisches Struct parsen und speichern
	switch projType {
	case "image":
		var img models.ImageProject
		if err := json.Unmarshal(jsonBody, &img); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image project"})
			return
		}
		if err := h.svc.Create(img); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image project"})
			return
		}
	case "movie":
		var vid models.MovieProject
		if err := json.Unmarshal(jsonBody, &vid); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video project"})
			return
		}
		if err := h.svc.Create(vid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save video project"})
			return
		}
	case "music":
		var mus models.MusicProject
		if err := json.Unmarshal(jsonBody, &mus); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid music project"})
			return
		}
		if err := h.svc.Create(mus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save music project"})
			return
		}
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
