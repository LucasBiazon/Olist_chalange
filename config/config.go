package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *Logger
)

func Init() error {
	var err error
	database, err = InitializeDatabase()
	if err != nil {
		return fmt.Errorf("Error initializing database: %v", err)
	}
	return nil
}
func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

func GetDatabase() *gorm.DB {
	return database
}
