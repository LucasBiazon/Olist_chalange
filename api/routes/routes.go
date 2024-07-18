package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api"
	v1 := router.Group(basePath)
	{
		v1.POST("/book", handler.CreateBook)
		// v1.GET("/book", handler.GetBook)
		v1.GET("/author", handler.GetAuthors)
	}
}
