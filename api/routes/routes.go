package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/handler"
	"github.com/lucasBiazon/olist/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.POST("/book", handler.CreateBookHandler)
		v1.GET("/book", handler.GetBookHandler)
		v1.PUT("/book", handler.UpdateBookHandler)
		v1.DELETE("/book", handler.DeleteBookHandler)
		v1.GET("/author", handler.GetAuthorHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
