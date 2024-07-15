package handler

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Author struct {
	Name string
}

func GetCSV() error {
	file, err := os.Open("./data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
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
