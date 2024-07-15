package handler

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Author struct {
	Name string
}

func GetCSV(filePath *os.File) error {

	records, err := csv.NewReader(filePath).ReadAll()
	if err != nil {
		return err
	}

	people := make([]Author, len(records))
	for index, record := range records {
		author := Author{
			Name: record[0],
		}
		people[index] = author
	}
	fmt.Printf("People: %v\n", people)
	return nil
}
