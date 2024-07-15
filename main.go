package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasBiazon/olist/APi/handler"
)

func CSVComand() {
	filePath := flag.String("csv", "", "File path")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("File path is required")
		fmt.Println("Please provide the path to the file using the -file flag.")
		os.Exit(1)
	}

	if !isCSVFile(*filePath) {
		fmt.Println("Invalid file format")
		fmt.Println("Please provide a CSV file")
		os.Exit(1)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("File %v opened successfully\n", *filePath)
	handler.GetCSV(file)

}

func isCSVFile(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	return ext == ".csv"
}

func main() {
	CSVComand()

}
