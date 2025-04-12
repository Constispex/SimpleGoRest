// Dieses Skript dient als Beispiel, wie Migrationsaufgaben gestartet werden können.
// In einer echten Umgebung würdest du hier ein Migrations-Tool wie "goose" oder "migrate" einbinden.
package main

import (
	"fmt"
	"prosting/backend-gin/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		fmt.Println("Could not load config:", err)
		return
	}
	fmt.Println("Running migrations for database:", cfg.Database)
	// Hier kommen die Migrationsbefehle, z. B.:
	// err = goose.Up(db, "./migrations")
	// if err != nil { ... }
}
