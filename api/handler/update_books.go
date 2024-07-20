package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

// UpdateBookHandler godoc
// @Summary Update a book
// @Description Update book details based on provided ID. Supports updating name, edition, publisher year, and authors.
// @Tags books
// @Accept  json
// @Produce  json
// @Param id query string true "Book ID"
// @Param book body types.UpdateBookRequest true "Update Book Request"
// @Success 200 {object} schema.Book
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /books [put]
func UpdateBookHandler(ctx *gin.Context) {
	request := &types.UpdateBookRequest{}
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := &schema.Book{}
	if err := database.Where("id = ?", id).First(&book).Error; err != nil {
		logger.Errorf("book not found: %v", err)
		types.SendError(ctx, http.StatusNotFound, "book not found")
		return
	}

	if request.Name != "" {
		book.Name = request.Name
	}
	if request.Edition != "" {
		book.Edition = request.Edition
	}
	if request.PublisherYear != "" {
		book.PublisherYear = request.PublisherYear
	}

	if len(request.Authors) > 0 {
		var authors []schema.Author
		for _, authorID := range request.Authors {
			var author schema.Author
			if err := database.Where("id = ?", authorID).First(&author).Error; err != nil {
				types.SendError(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			authors = append(authors, author)
		}

		if err := database.Model(&book).Association("Authors").Replace(authors); err != nil {
			types.SendError(ctx, http.StatusInternalServerError, "failed to update authors association")
			return
		}

		for _, author := range authors {
			if err := database.Model(&author).Association("Books").Append(book); err != nil {
				types.SendError(ctx, http.StatusInternalServerError, "failed to update books association for author")
				return
			}
		}
	}

	if err := database.Save(&book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookResponse := &schema.Book{
		ID:            book.ID,
		Name:          book.Name,
		Edition:       book.Edition,
		PublisherYear: book.PublisherYear,
		Authors:       book.Authors,
	}
	types.SendSuccess(ctx, "update-book", bookResponse)
}
