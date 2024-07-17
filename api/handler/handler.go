package handler

import (
	"github.com/lucasBiazon/olist/config"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *config.Logger
)

func InitializeHandler() {
	database = config.GetDatabase()
	logger = config.GetLogger("Handler")
}
