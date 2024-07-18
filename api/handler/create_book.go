package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

func CreateBook(ctx *gin.Context) {
	request := &types.CreateBookRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		types.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		types.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	authors := []*schema.Author{}
	for _, id := range request.Authors {
		var author schema.Author
		if err := database.Where("ID = ?", id).Find(&authors).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find authors"})
			return
		}
		authors = append(authors, &author)
	}

	id := uuid.New().String()
	book := schema.Book{
		ID:            id,
		Name:          request.Name,
		Edition:       request.Edition,
		PublisherYear: request.PublisherYear,
		Authors:       authors,
	}

	if err := database.Create(&book).Error; err != nil {
		types.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	types.SendSuccess(ctx, "CreateBook", book)

}
