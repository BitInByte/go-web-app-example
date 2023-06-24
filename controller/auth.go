package controller

import (
	"fmt"
	"net/http"

	"github.com/BitInByte/web-app-example/core"
	"github.com/BitInByte/web-app-example/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupBodyDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type LoginBodyDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
}

func AuthSignup(ctx *gin.Context) {
	var signupBody SignupBodyDTO

	// if ctx.BindJSON(&signupBody) != nil {
	// Fetch body data and validate
	if ctx.Bind(&signupBody) != nil {
		fmt.Println(signupBody)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse body. Please forward all data!",
		})
		return
	}
	// fmt.Println(signupBody)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupBody.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong hashing the password!",
		})
		return
	}

	// Create the use on database
	user := model.User{Username: signupBody.Email, Email: signupBody.Email, Password: string(hashedPassword)}
	result := core.DB.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong storing the user on database!",
		})
		return

	}

	// Empty password
	signupBody.Password = ""

	// Return response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created with success",
		"data":    signupBody,
	})

}

func AuthLogin(ctx *gin.Context) {
	var loginBody LoginBodyDTO

	// Fetch body data and validate
	if ctx.Bind(&loginBody) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse body. Please forward all data!",
		})
		return
	}

	var user model.User
	// Populates data to the user struct
	result := core.DB.First(&user, "email = ?", loginBody.Email)
	// SELECT * FROM users WHERE email = {Email};
	fmt.Println(result.RowsAffected, result.Error)
	if result.RowsAffected == 0 || result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch user from database! Email might be invalid.",
		})
		return
	}

	// Check password with stored encrypted password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password)) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Password",
		})
		return

	}

	// Send response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Login successfuly",
	})
}
