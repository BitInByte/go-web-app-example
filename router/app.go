package router

import "github.com/gin-gonic/gin"

func AppRouter(router gin.RouterGroup)  {
    router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test",
		})
    })
}
