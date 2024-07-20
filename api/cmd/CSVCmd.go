package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/handler"
	"github.com/lucasBiazon/olist/config"
)

var (
	logger *config.Logger
)

func CSVCmd(ctx *gin.Context) {
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

	err = handler.CreateAuthorHandler(file)
	if err != nil {
		logger.Errorf("Error creating authors: %v", err)
		os.Exit(1)
	}

}

func isCSVFile(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	return ext == ".csv"
}

func main() {
	var ctx *gin.Context
	logger := config.GetLogger("Main")
	if err := config.Init(); err != nil {
		logger.Errorf("Error initializing application: %v", err)
	}
	CSVCmd(ctx)
}
