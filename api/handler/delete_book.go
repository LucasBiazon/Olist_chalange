package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

// DeleteBookHandler godoc
// @Summary Delete a book
// @Description Deletes a book from the database based on the provided ID.
// @Tags books
// @Accept  json
// @Produce  json
// @Param id query string true "Book ID"
// @Success 200 {object} schema.Book
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /books [delete]
func DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	book := &schema.Book{}

	if err := database.Where("id = ?", id).First(&book).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "book not found"})
		return
	}

	if err := database.Where("id = ?", id).Delete(&book).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	types.SendSuccess(ctx, "delete-book", book)

}
