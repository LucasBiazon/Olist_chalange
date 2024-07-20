package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

// CreateBookHandler godoc
// @Summary Create a new book
// @Description Creates a new book with the provided details and associates authors to it.
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body types.CreateBookRequest true "Create Book Request"
// @Success 201 {object} schema.Book
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /books [post]

func CreateBookHandler(ctx *gin.Context) {
	request := &types.CreateBookRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		types.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		types.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := uuid.NewString()
	book := schema.Book{
		ID:            id,
		Name:          request.Name,
		Edition:       request.Edition,
		PublisherYear: request.PublisherYear,
	}

	if err := database.Create(&book).Error; err != nil {
		types.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	for _, authorID := range request.Authors {
		var author schema.Author
		if err := database.First(&author, "ID = ?", authorID).Error; err != nil {
			types.SendError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := database.Model(&book).Association("Authors").Append(&author); err != nil {
			types.SendError(ctx, http.StatusInternalServerError, "f")
			return
		}

		if err := database.Model(&author).Association("Books").Append(&book); err != nil {
			types.SendError(ctx, http.StatusInternalServerError, "a")
			return
		}
	}

	types.SendSuccess(ctx, "201", book)
}
