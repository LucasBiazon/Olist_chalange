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
		v1.POST("/book", handler.CreateBookHandler)
		v1.GET("/book", handler.GetBookHandler)
		v1.PUT("/book", handler.UpdateBookHandler)
		v1.DELETE("/book", handler.DeleteBookHandler)
		// v1.POST("/author", handler.CreateAuthorHandler)
		v1.GET("/author", handler.GetAuthorHandler)
	}
}
