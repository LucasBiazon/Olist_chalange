package config

import (
	"os"

	"github.com/lucasBiazon/olist/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("Sqlite")
	dbPatch := "./database/database.db"
	_, err := os.Stat(dbPatch)

	if os.IsNotExist(err) {
		logger.Info("Database not found. Creating new database")
		err := os.MkdirAll("./database", os.ModePerm)
		if err != nil {
			logger.Error("Error creating database directory")
			return nil, err
		}
		file, err := os.Create(dbPatch)
		if err != nil {
			logger.Error("Error creating database file")
			return nil, err
		}
		file.Close()
	}

	database, err := gorm.Open(sqlite.Open(dbPatch), &gorm.Config{})
	if err != nil {
		logger.Error("Error opening database")
		return nil, err
	}
	// Create database tables
	if err := database.AutoMigrate(&schema.Book{}, &schema.Author{}); err != nil {
		logger.Errorf("Error migrating database schema: %v", err)
		return nil, err
	}
	logger.Info("Database initialized successfully")
	return database, nil
}
