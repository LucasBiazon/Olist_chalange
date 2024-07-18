package handler

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/lucasBiazon/olist/schema"
)

func CreateAuthors(filePath *os.File) error {

	InitializeHandler()
	records, err := csv.NewReader(filePath).ReadAll()
	if err != nil {
		return err
	}
	if database == nil {
		return fmt.Errorf("database is not initialized")
	}
	logger.Infof("Creating authors from file: %s", filePath.Name())
	for _, record := range records {
		author := schema.Author{
			ID:   uuid.NewString(),
			Name: record[0],
		}
		if err := database.Create(&author).Error; err != nil {
			logger.Errorf("failed to create author: %v", err)
			return nil
		}
		logger.Infof("Author created: UUID=%s, Nome=%s", author.ID, author.Name)

	}
	return nil
}
