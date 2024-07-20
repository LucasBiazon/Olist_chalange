package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

// GetBookHandler godoc
// @Summary Get books by various filters
// @Description Get books based on provided filters (id, name, edition, publisher_year) or return all books if no filter is provided
// @Tags books
// @Accept  json
// @Produce  json
// @Param id query string false "Book ID"
// @Param name query string false "Book Name"
// @Param edition query string false "Book Edition"
// @Param publisher_year query string false "Publisher Year"
// @Success 200 {array} map[string]interface{}
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /books [get]
func GetBookHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")
	edition := ctx.Query("edition")
	publisherYear := ctx.Query("publisher_year")

	if id == "" && name == "" && edition == "" && publisherYear == "" {
		GetBooksHandler(ctx)
		return
	}

	books := []*schema.Book{}
	query := database.Model(&schema.Book{})

	if id != "" {
		query = query.Where("id = ?", id)
		if name != "" {
			query = query.Where("name = ?", name, "id = ?", id)
		}
		if edition != "" {
			query = query.Where("edition = ?", edition, "id = ?", id)
		}
		if publisherYear != "" {
			query = query.Where("publisher_year = ?", publisherYear, "id = ?", id)
		}
	} else if name != "" {
		query = query.Where("name = ?", name)
		if edition != "" {
			query = query.Where("edition = ?", edition, "name = ?", name)
		}
		if publisherYear != "" {
			query = query.Where("publisher_year = ?", publisherYear, "name = ?", name)
		}
	} else if edition != "" {
		query = query.Where("edition = ?", edition)
		if publisherYear != "" {
			query = query.Where("publisher_year = ?", publisherYear, "edition = ?", edition)
		}
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

	types.SendSuccess(ctx, "get-book(s) response", response)
}
