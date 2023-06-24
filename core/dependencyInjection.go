package core

import (
	// "fmt"

	"github.com/BitInByte/web-app-example/controller"
	"github.com/BitInByte/web-app-example/middleware"
	"github.com/BitInByte/web-app-example/router"
	"gorm.io/gorm"
)

type DependencyInjectionContainer struct {
	DB *gorm.DB

	AuthRouter *router.AuthRouter
	TodoRouter *router.TodoRouter

	AuthController *controller.AuthController
	TodoController *controller.TodoController

	AuthMiddleware *middleware.AuthMiddleware
}

func NewDependencyInjectionContainer() *DependencyInjectionContainer {
	dic := &DependencyInjectionContainer{
		DB: LoadSqliteDBSettings(),
	}

	// Order here matters. We should always start instantiate the
	// last dependencies that are needed and the first dependencies
	// needed to the end
	// Instantiate
	// controllers -> middleware -> router
	// Execution
	// router -> middleware -> controllers

	// Controllers
	{
		dic.AuthController = &controller.AuthController{DB: dic.DB}
		dic.TodoController = &controller.TodoController{DB: dic.DB}
	}

	// Middleware
	{
		dic.AuthMiddleware = &middleware.AuthMiddleware{DB: dic.DB}
	}

	// Routers
	{
		dic.AuthRouter = &router.AuthRouter{AuthController: dic.AuthController, AuthMiddleware: dic.AuthMiddleware}
		dic.TodoRouter = &router.TodoRouter{TodoController: dic.TodoController}
	}

	return dic
}
