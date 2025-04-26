package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"prosting/backend-gin/internal/handlers"
	"prosting/backend-gin/internal/repository"
	"prosting/backend-gin/internal/services"
	"prosting/backend-gin/scripts"
	"time"
)

func main() {
	// Datenbankverbindung herstellen
	scripts.RunMigration()
	//Gin-Router initialisieren
	r := gin.Default()

	r.MaxMultipartMemory = 1024 << 20
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS-Konfiguration
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	// Initialisiere den Projekt-Repository und -Service
	repo := repository.NewProjectRepository(scripts.DB)
	svc := services.NewProjectService(repo)
	h := handlers.NewProjectHandler(svc)

	r.Static("/uploads", "./uploads")

	// Routen definieren
	r.POST("/api/projects", h.Create)
	r.GET("/api/projects", h.GetAll)
	r.GET("/api/categories", h.GetCategories)

	//Starte den Server auf dem angegebenen Port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server startet auf Port %s", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
		return
	}
}

func init() {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}
}
