package scripts

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"prosting/backend-gin/pkg/config"
)

func RunMigration() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		fmt.Println("Could not load config:", err)
		return
	}
	fmt.Println("Running migrations for database:", cfg.Database)

	db, err := sql.Open("postgres", getConnectionString(cfg))
	if err != nil {
		fmt.Println("Could not connect to database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		_ = fmt.Errorf("could not ping database: %s", err)
	}
	defer db.Close()
	fmt.Println("Connected to database successfully")

}

func getConnectionString(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.Database)
}
