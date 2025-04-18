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
	router := gin.Default()

	// User-Repository und -Service initialisieren
	userRepo := repository.NewUserRepository(scripts.DB)
	userService := services.NewUserService(userRepo)

	// User-Handler erstellen und Routen definieren
	userHandler := handlers.NewUserHandler(userService)

	// Definiere die Routen
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)

	// Starte den Server auf dem angegebenen Port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server startet auf Port %s", port)
	router.Run(":" + port)
}
