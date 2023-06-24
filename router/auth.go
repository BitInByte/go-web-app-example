package router

import (
	"github.com/BitInByte/web-app-example/controller"
	"github.com/BitInByte/web-app-example/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) {
	router.POST("/signup", controller.AuthSignup)
	router.POST("/login", controller.AuthLogin)
	router.GET("/", middleware.AuthGuard, controller.Validate)
}
