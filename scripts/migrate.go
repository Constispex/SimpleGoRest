package scripts

import (
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"prosting/backend-gin/internal/models"
	"prosting/backend-gin/pkg/config"
)

var DB *gorm.DB

func RunMigration() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		fmt.Println("Could not load config:", err)
		return
	}
	fmt.Println("Running migrations for database:", cfg.Database)

	db, err := gorm.Open(postgres.Open(getConnectionString(cfg)), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to database:", err)
	}

	// Try to Ping the database
	if err := db.Exec("SELECT 1").Error; err != nil {
		fmt.Println("Could not ping database:", err)
		return
	}
	DB = db

	// Migrate the schema
	if err := db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Room{},
		&models.Category{},
	); err != nil {
		fmt.Println("Could not migrate database:", err)
		return
	}

	fmt.Println("Database migration completed successfully.")

}

func getConnectionString(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.Database)
}
