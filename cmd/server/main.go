package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"prosting/backend-gin/internal/handlers"
	"prosting/backend-gin/internal/repository"
	"prosting/backend-gin/internal/services"
	"prosting/backend-gin/scripts"
)

func main() {
	// Datenbankverbindung herstellen
	scripts.RunMigration()
	//Gin-Router initialisieren
	r := gin.Default()
	repo := repository.NewProjectRepository(scripts.DB)
	svc := services.NewProjectService(repo)
	h := handlers.NewProjectHandler(svc)

	// Routen definieren
	r.POST("/api/projects", h.Create)
	r.GET("/api/projects", h.GetAll)

	//Starte den Server auf dem angegebenen Port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server startet auf Port %s", port)
	r.Run(":" + port)
}
