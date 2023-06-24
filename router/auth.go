package router

import (
	"fmt"

	"github.com/BitInByte/web-app-example/controller"
	"github.com/BitInByte/web-app-example/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type AuthRouter struct {
	DB             *gorm.DB
	AuthController controller.AuthController
	AuthMiddleware middleware.AuthMiddleware
	// DI core.DependencyInjectionContainer
}

func (a AuthRouter) Init(router *gin.RouterGroup) {
	// a.authController = controller.AuthController{
	// 	DB: a.DB,
	// }
	//
	// a.authMiddleware = middleware.AuthMiddleware{
	// 	DB: a.DB,
	// }
	fmt.Println("Auth Router", a.DB == nil, a.AuthController.DB == nil)

	router.POST("/signup", a.AuthController.AuthSignup)
	router.POST("/login", a.AuthController.AuthLogin)
	router.GET("/", a.AuthMiddleware.AuthGuard, a.AuthController.Validate)
}
