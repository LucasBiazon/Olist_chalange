package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/schema"
)

func GetAuthors(ctx *gin.Context) {

	authors := []schema.Author{}
	if err := database.Preload("Books").Find(&authors).Error; err != nil {
		logger.Errorf("failed to get authors: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get authors"})
		return
	}
	var responseAuthors []gin.H
	for _, author := range authors {
		responseAuthors = append(responseAuthors, gin.H{
			"ID":        author.ID,
			"Name":      author.Name,
			"CreatedAt": author.CreatedAt,
			"UpdatedAt": author.UpdatedAt,
			"DeletedAt": author.DeletedAt,
			"Books":     author.Books, // Isso deve conter os livros associados
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"authors": responseAuthors})
}
