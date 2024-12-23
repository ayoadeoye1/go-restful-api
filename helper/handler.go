package helper

import (
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := responses.Success(statusCode, message, data)
	ctx.JSON(statusCode, response)
}

func SendError(ctx *gin.Context, statusCode int, message string, err interface{}) {
	response := responses.Error(statusCode, message, err)
	ctx.JSON(statusCode, response)
}
