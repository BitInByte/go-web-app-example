package controller

import (
	"fmt"
	"net/http"

	"github.com/BitInByte/web-app-example/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

type createTodoBodyDTO struct {
	// Title string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}

func (t TodoController) CreateTodo(ctx *gin.Context) {
	var createTodoBody createTodoBodyDTO
	// Get data from request body
	if ctx.Bind(&createTodoBody) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse body. Please forward all data!",
		})
		return
	}

	// Get logged in user from request
	user, _ := ctx.Get("user")
	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get session user!",
		})
		return

	}
	// fmt.Print("User", user)

	userId := user.(model.User).ID
	// Store todo
	todo := model.Todo{
		// Title:  createTodoBody.Title,
		Body:   createTodoBody.Body,
		UserID: userId,
		Status: model.Created,
	}
	result := t.DB.Create(&todo)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong storing the todo on database!",
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Todo created with success",
		"data":    todo,
	})
}

func (t TodoController) GetAllTodos(ctx *gin.Context) {
	// Get logged in user from session
	user, _ := ctx.Get("user")
	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get session user!",
		})
		return

	}
	userId := user.(model.User).ID

	// Get todos from user
	var foundUser model.User
	result := t.DB.Model(model.User{}).Preload("Todos").First(&foundUser, "id = ?", userId)
	if result == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get todos!",
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todos retrieved with success",
		"data":    foundUser.Todos,
	})

}

func (t TodoController) GetAllStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Status retrieved with success",
		"data": []string{
			"done",
			"in progress",
			"created",
		},
	})
}

type StatusUriDTO struct {
	ID string `uri:"id" binding:"required"`
}

// type StatusBodyDTO struct {
// 	Status string `uri:"status" binding:"required"`
// }

func (t TodoController) ChangeStatus(ctx *gin.Context) {
	var statusUriDTO StatusUriDTO
	if ctx.BindUri(&statusUriDTO) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing id from the URL",
		})
		return
	}
	// or
	// fmt.Println(ctx.Param("id"))

	// var statusBodyDTO StatusBodyDTO
	// if ctx.Bind(&statusBodyDTO) != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Missing the current status from the body",
	// 	})
	// 	return
	// }

	fmt.Println(statusUriDTO)

	// Get todo from id
	var todo model.Todo
	t.DB.First(&todo, statusUriDTO.ID)

	switch todo.Status {
	case "created":
		todo.Status = "in progress"
	case "in progress":
		todo.Status = "done"
	}

	t.DB.Save(&todo)

	// Could also be done like this:
	// t.DB.Model(&todo).Updates(model.Todo{
	//     Body: todo.Body,
	//     Status: "status here",
	// })

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Status retrieved with success",
		"data":    todo,
	})
}

func (t TodoController) DeleteTodo(ctx *gin.Context) {
	var deleteUriDTO StatusUriDTO
	if ctx.BindUri(&deleteUriDTO) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing id from the URL",
		})
		return
	}

	fmt.Println(deleteUriDTO.ID)
	t.DB.Delete(&model.Todo{}, &deleteUriDTO.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfuly deleted",
		"data":    deleteUriDTO.ID,
	})
}
