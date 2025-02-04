package migrations

import (
	"log"

	"github.com/adimurianto/go-gin-boilerplate/infra/database"
	"github.com/adimurianto/go-gin-boilerplate/models"
)

// Migrate Add list of model add for migrations
func Migrate() {
	// Enable UUID extension
	if err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Printf("Failed to create UUID extension: %v", err)
		return
	}

	var migrationModels = []interface{}{
		&models.User{},
		// add other models here
	}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		log.Printf("Migration error: %v", err)
		return
	}
}
