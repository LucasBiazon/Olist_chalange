package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/config"
)

var (
	logger config.Logger
)

func main() {
	route := gin.Default()
	logger := config.GetLogger("Main")
	if err := config.Init(); err != nil {
		logger.Errorf("Error initializing application: %v", err)
	}
	route.Run(":8080")
}
