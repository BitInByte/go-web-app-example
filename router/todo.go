package router

import (
	"github.com/BitInByte/web-app-example/controller"
	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
	TodoController *controller.TodoController
}

func (t TodoRouter) Init(router *gin.RouterGroup) {
	router.GET("/", t.TodoController.GetAllTodos)
	router.POST("/", t.TodoController.CreateTodo)
}
