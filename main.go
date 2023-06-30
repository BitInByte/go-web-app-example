package main

import (
	"github.com/BitInByte/web-app-example/core"
	"github.com/BitInByte/web-app-example/router"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var di *core.DependencyInjectionContainer

// Runs before main
// Perfect for initializations
func init() {
	core.LoadEnvVariables()
	di = core.NewDependencyInjectionContainer()
	core.Migrations(di.DB)
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")

	r.Use(static.Serve("/", static.LocalFile("./view/dist", true)))

	{
		app := v1.Group("/")
		router.AppRouter(app)

		auth := v1.Group("/auth")
		di.AuthRouter.Init(auth)

		todo := v1.Group("/todo", di.AuthMiddleware.AuthGuard)
		di.TodoRouter.Init(todo)
	}

	r.Run()
}
