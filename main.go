package main

import (
	"github.com/lucasBiazon/olist/config"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("Main")
	if err := config.Init(); err != nil {
		logger.Errorf("Error initializing application: %v", err)
	}
}
