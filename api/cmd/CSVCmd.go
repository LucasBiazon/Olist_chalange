package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasBiazon/olist/api/handler"
	"github.com/lucasBiazon/olist/config"
)

var (
	logger *config.Logger
)

func CSVCmd() {
	csvPath := flag.String("csv", "", "File path")
	flag.Parse()

	if *csvPath == "" {
		logger.Error("File path is required")
		logger.Warning("Please provide the path to the file using the -csv flag.")
		logger.Warning("Example: go run main.go -csv /path/to/file.csv")
		os.Exit(1)
	}

	if !isCSVFile(*csvPath) {
		logger.Error("Invalid file format, please provide a CSV file")
		os.Exit(1)
	}

	file, err := os.Open(*csvPath)
	if err != nil {
		logger.Errorf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("CSV file: %v opened successfully\n", *csvPath)
	handler.GetCSV(file)
}

func isCSVFile(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	return ext == ".csv"
}

func main() {
	CSVCmd()
}
