package main

import (
	"log"

	"github.com/BitInByte/web-app-example/core"
	"github.com/BitInByte/web-app-example/router"
	"github.com/gin-gonic/gin"
)

func init() {
    core.LoadEnvVariables()
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

    log.Println("App listening on port 8000")
	r.Run() // listen and serve on 0.0.0.0:8080
}
