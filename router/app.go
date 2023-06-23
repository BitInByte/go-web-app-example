package router

import (
	"github.com/BitInByte/web-app-example/controller"
	"github.com/gin-gonic/gin"
)

func AppRouter(router *gin.RouterGroup)  {
    router.GET("/", controller.AppView)
}
