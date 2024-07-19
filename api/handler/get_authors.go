package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/schema"
)

func GetAuthors(ctx *gin.Context) {

	authors := []schema.Author{}
	if err := database.Preload("Books.Authors").Find(&authors).Error; err != nil {
		logger.Errorf("failed to get authors: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get authors"})
		return
	}

	var responseAuthors []gin.H
	for _, author := range authors {
		var responseBooks []gin.H
		for _, book := range author.Books {
			var responseBookAuthors []gin.H
			for _, bookAuthor := range book.Authors {
				responseBookAuthors = append(responseBookAuthors, gin.H{
					"ID":   bookAuthor.ID,
					"Name": bookAuthor.Name,
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
		responseAuthors = append(responseAuthors, gin.H{
			"ID":        author.ID,
			"Name":      author.Name,
			"CreatedAt": author.CreatedAt,
			"UpdatedAt": author.UpdatedAt,
			"DeletedAt": author.DeletedAt,
			"Books":     responseBooks,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"authors": responseAuthors})
}
