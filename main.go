package main

import (
	"log"

	"github.com/BitInByte/web-app-example/router"
	"github.com/gin-gonic/gin"
)

func main()  {
    r := gin.Default()
    v1 := r.Group("/v1")
    {
        app := v1.Group("/")
        router.AppRouter(*app)
    }

    log.Println("App listening on port 8000")
	r.Run() // listen and serve on 0.0.0.0:8080
}
