package handler

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/lucasBiazon/olist/api/schema"
)

func CreateAuthors(filePath *os.File) error {
	InitializeHandler()
	records, err := csv.NewReader(filePath).ReadAll()
	if err != nil {
		return err
	}
	if database == nil {
		return fmt.Errorf("database não inicializado corretamente")
	}
	for _, record := range records {
		id := uuid.New().String()
		author := schema.Author{
			ID:   id,
			Name: record[0],
		}
		if err := database.Create(&author).Error; err != nil {
			logger.Errorf("Error creating author: %v", err.Error())
			return nil
		}
		fmt.Printf("Author created: %v\n", author)
	}
	return nil
}
