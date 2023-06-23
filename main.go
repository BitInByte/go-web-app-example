package main

import (
	"github.com/BitInByte/web-app-example/core"
	"github.com/BitInByte/web-app-example/router"
	"github.com/gin-gonic/gin"
)

// Runs before main
// Perfect for initializations
func init() {
    core.LoadEnvVariables()
    core.LoadDBSettings()
    core.Migrations()
}

func main()  {
    r := gin.Default()
    v1 := r.Group("/v1")

    {
        app := v1.Group("/")
        router.AppRouter(app)

        auth := v1.Group("/auth")
        router.AuthRouter(auth)
    }

	r.Run()
}
