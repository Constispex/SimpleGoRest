package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort  string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DefEmail    string
	DefPassword string
	Database    string
}

func LoadConfig(filename string) (Config, error) {
	// Lade die .env-Datei
	err := godotenv.Load(filename)
	if err != nil {
		return Config{}, err
	}

	// Konfiguration aus Umgebungsvariablen lesen
	return Config{
		ServerPort:  os.Getenv("SERVER_PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("POSTGRES_USER"),
		DBPass:      os.Getenv("POSTGRES_PASSWORD"),
		DefEmail:    os.Getenv("PGADMIN_DEFAULT_EMAIL"),
		DefPassword: os.Getenv("PGADMIN_DEFAULT_PASSWORD"),
		Database:    os.Getenv("POSTGRES_DB"),
	}, nil
}
