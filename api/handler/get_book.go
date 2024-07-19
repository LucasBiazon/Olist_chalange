package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/schema"
)

func GetBook(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")
	edition := ctx.Query("edition")
	publisherYear := ctx.Query("publisher_year")

	if id == "" && name == "" && edition == "" && publisherYear == "" {
		GetBooks(ctx)
	}

	books := []*schema.Book{}
	query := database.Model(&schema.Book{})

	if id != "" {
		query = query.Where("id = ?", id)
	} else if name != "" {
		query = query.Where("name = ?", name)
	} else if edition != "" {
		query = query.Where("edition = ?", edition)
	} else if publisherYear != "" {
		query = query.Where("publisher_year = ?", publisherYear)
	}

	if err := query.Find(&books).Error; err != nil {
		logger.Errorf("book not found: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	response := []map[string]interface{}{}
	for _, book := range books {
		response = append(response, map[string]interface{}{
			"id":             book.ID,
			"name":           book.Name,
			"edition":        book.Edition,
			"publisher_year": book.PublisherYear,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
