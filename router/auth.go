package router

import (
	"github.com/BitInByte/web-app-example/controller"
	"github.com/BitInByte/web-app-example/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	AuthController *controller.AuthController
	AuthMiddleware *middleware.AuthMiddleware
}

func (a AuthRouter) Init(router *gin.RouterGroup) {
	router.POST("/signup", a.AuthController.AuthSignup)
	router.POST("/login", a.AuthController.AuthLogin)
	router.GET("/", a.AuthMiddleware.AuthGuard, a.AuthController.Validate)
}
