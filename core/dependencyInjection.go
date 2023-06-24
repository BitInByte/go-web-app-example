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

	AuthRouter router.AuthRouter
	TodoRouter router.TodoRouter

	AuthController controller.AuthController
	TodoController controller.TodoController

	AuthMiddleware middleware.AuthMiddleware
}

// func (d DependencyInjectionContainer) LoadDependencyInjectionContainer() {
// 	d.DB = LoadSqliteDBSettings()
// }

func NewDependencyInjectionContainer() *DependencyInjectionContainer {
	dic := DependencyInjectionContainer{
		DB: LoadSqliteDBSettings(),
	}
	// newDependencyInjectionContainer.LoadDependencyInjectionContainer()
	db := LoadSqliteDBSettings()

	// Routers
	// {
	dic.AuthRouter = router.AuthRouter{DB: dic.DB, AuthController: dic.AuthController, AuthMiddleware: dic.AuthMiddleware}
	dic.TodoRouter = router.TodoRouter{TodoController: dic.TodoController}
	// }

	// Controllers
	// {
	dic.AuthController = controller.AuthController{DB: dic.DB}
	dic.TodoController = controller.TodoController{DB: db}
	// }

	// Middleware
	// {
	dic.AuthMiddleware = middleware.AuthMiddleware{DB: db}
	// }

	// dic.DB = db
	// fmt.Println(DB)
	return &dic
}
