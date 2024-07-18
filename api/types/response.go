package types

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/olist/schema"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"error":     msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateBookResponse struct {
	Message string      `json:"message"`
	Data    schema.Book `json:"data"`
}

type DeleteBookResponse struct {
	Message string      `json:"message"`
	Data    schema.Book `json:"data"`
}
type ShowBookResponse struct {
	Message string      `json:"message"`
	Data    schema.Book `json:"data"`
}
type ListBookResponse struct {
	Message string      `json:"message"`
	Data    schema.Book `json:"data"`
}
type UpdateBoookResponse struct {
	Message string      `json:"message"`
	Data    schema.Book `json:"data"`
}
