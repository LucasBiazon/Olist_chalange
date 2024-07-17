package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lucasBiazon/olist/api/schema"
	"github.com/lucasBiazon/olist/api/types"
)

func CreatBook(ctx *gin.Context) {
	request := &types.CreateBookRequest{}
	ctx.Bind(request)
	if err := request.Validate(); err != nil {
		return
	}
	id := uuid.New().String()
	book := &schema.Book{
		ID:            id,
		Name:          request.Name,
		Edition:       request.Edition,
		PublisherYear: request.PublisherYear,
		Authors:       request.AuthorId,
	}

	if err := database.Create(book).Error; err != nil {
		logger.Errorf("Error creating book: %v", err.Error())
		return
	}
	ctx.JSON(200, book)
}
