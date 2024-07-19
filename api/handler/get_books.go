package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/schema"
)

func GetBooks(ctx *gin.Context) {
	books := []schema.Book{}
	if err := database.Preload("Authors").Find(&books).Error; err != nil {
		logger.Errorf("failed to get authors: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get authors"})
		return
	}

	var responseBooks []gin.H
	for _, book := range books {
		var responseBookAuthors []gin.H
		for _, author := range book.Authors {
			responseBookAuthors = append(responseBookAuthors, gin.H{
				"ID":   author.ID,
				"Name": author.Name,
			})
		}
		responseBooks = append(responseBooks, gin.H{
			"ID":            book.ID,
			"Name":          book.Name,
			"Edition":       book.Edition,
			"PublisherYear": book.PublisherYear,
			"CreatedAt":     book.CreatedAt,
			"UpdatedAt":     book.UpdatedAt,
			"DeletedAt":     book.DeletedAt,
			"Authors":       responseBookAuthors,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"Books": responseBooks})
}
