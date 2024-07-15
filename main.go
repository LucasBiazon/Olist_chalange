package main

import (
	"fmt"

	"github.com/lucasBiazon/olist/APi/handler"
)

func main() {
	err := handler.GetCSV()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
