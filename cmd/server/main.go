package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"prosting/backend-gin/internal/handlers"
	"prosting/backend-gin/pkg/config"
	"prosting/backend-gin/scripts"
)

func main() {
	// Konfiguration laden (z. B. Port, DB-Details aus .env)
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Datenbankverbindung herstellen
	scripts.RunMigration()
	//Gin-Router initialisieren
	router := gin.Default()

	// User-Handler erstellen und Routen definieren
	userHandler := handlers.NewUserHandler()
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)

	// Starte den Server auf dem angegebenen Port
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.DBPort
	}
	log.Printf("Server startet auf Port %s", port)
	router.Run(":" + port)
}
