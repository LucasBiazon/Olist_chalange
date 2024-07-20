package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/api/types"
	"github.com/lucasBiazon/olist/schema"
)

// GetAuthorHandler godoc
// @Summary Get authors by various filters
// @Description Get authors based on provided filters (name) or return all authors with pagination
// @Tags authors
// @Accept  json
// @Produce  json
// @Param name query string false "Author Name"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of authors per page" default(10)
// @Success 200 {object} gin.H
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /authors [get]

func GetAuthorHandler(ctx *gin.Context) {

	name := ctx.Query("name")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	authors := []*schema.Author{}
	var responseAuthors []gin.H
	var total int64

	if name != "" {
		if err := database.Model(&schema.Author{}).Where("Name = ?", name).Count(&total).Error; err != nil {
			logger.Errorf("failed to count authors: %v", err)
			types.SendError(ctx, http.StatusInternalServerError, "failed to count authors")
			return
		}

		if err := database.Where("Name = ?", name).Preload("Books.Authors").Offset(offset).Limit(pageSize).Find(&authors).Error; err != nil {
			logger.Errorf("author not found: %v", err)
			types.SendError(ctx, http.StatusNotFound, "author not found")
			return
		}

		if len(authors) == 0 {
			types.SendError(ctx, http.StatusNotFound, "no authors found")
			return
		}
		responseAuthors = buildAuthorsResponse(authors)
		types.SendSuccess(ctx, "get author(s)", gin.H{
			"authors":    responseAuthors,
			"page":       page,
			"pageSize":   pageSize,
			"total":      total,
			"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
		})
		return
	}
	if err := database.Model(&schema.Author{}).Count(&total).Error; err != nil {
		logger.Errorf("failed to count authors: %v", err)
		types.SendError(ctx, http.StatusInternalServerError, "failed to count authors")
		return
	}

	if err := database.Preload("Books.Authors").Offset(offset).Limit(pageSize).Find(&authors).Error; err != nil {
		logger.Errorf("failed to get authors: %v", err)
		types.SendError(ctx, http.StatusInternalServerError, "failed to get authors")
		return
	}

	responseAuthors = buildAuthorsResponse(authors)
	types.SendSuccess(ctx, "get author(s)", gin.H{
		"authors":    responseAuthors,
		"page":       page,
		"pageSize":   pageSize,
		"total":      total,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

func buildAuthorsResponse(authors []*schema.Author) []gin.H {
	var responseAuthors []gin.H
	for _, author := range authors {
		var responseBooks []string
		for _, book := range author.Books {
			responseBooks = append(responseBooks, book.ID)
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
	return responseAuthors
}
