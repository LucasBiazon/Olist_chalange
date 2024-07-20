package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

func GetBooksHandler(ctx *gin.Context) {
	books := []*schema.Book{}
	if err := database.Preload("Authors").Find(&books).Error; err != nil {
		logger.Errorf("failed to get authors: %v", err)
		types.SendError(ctx, http.StatusInternalServerError, "failed to get authors")
		return
	}

	var responseBooks []gin.H
	for _, book := range books {
		var responseBookAuthors []string
		for _, author := range book.Authors {
			responseBookAuthors = append(responseBookAuthors, author.ID)
		}
		responseBooks = append(responseBooks, gin.H{
			"ID":            book.ID,
			"Name":          book.Name,
			"Edition":       book.Edition,
			"PublisherYear": book.PublisherYear,
			"Authors":       responseBookAuthors,
		})
	}

	types.SendSuccess(ctx, "get book(s)", responseBooks)
}
