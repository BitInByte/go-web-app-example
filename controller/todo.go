package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func (t TodoController) CreateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Todo created with success",
		// "data":    signupBody,
	})
}
