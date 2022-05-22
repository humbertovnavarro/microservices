package apiHelpers

import "github.com/gin-gonic/gin"

func ErrorJSON(ctx *gin.Context, code uint32, message string) {
	ctx.JSON(int(code), gin.H{
		"error": message,
	})
}
