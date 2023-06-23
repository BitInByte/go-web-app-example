package controller

import "github.com/gin-gonic/gin"

func AppView(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "test",
	})
}
