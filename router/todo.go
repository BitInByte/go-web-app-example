package router

import (
	"github.com/BitInByte/web-app-example/controller"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type TodoRouter struct {
	// DB             *gorm.DB
	TodoController controller.TodoController
}

func (t TodoRouter) Init(router *gin.RouterGroup) {
	// t.todoController = controller.TodoController{DB: t.DB}
	router.POST("/", t.TodoController.CreateTodo)
}
