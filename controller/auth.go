package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	// "github.com/BitInByte/web-app-example/core"
	"github.com/BitInByte/web-app-example/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

type AuthController struct {
	DB *gorm.DB
}

func (a AuthController) AuthSignup(ctx *gin.Context) {
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
	// result := core.DB.Create(&user) // pass pointer of data to Create
	result := a.DB.Create(&user) // pass pointer of data to Create
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

func (a AuthController) AuthLogin(ctx *gin.Context) {
	fmt.Println("Auth Controller", a.DB == nil)
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
	// result := core.DB.First(&user, "email = ?", loginBody.Email)
	fmt.Println(a.DB == nil)
	result := a.DB.First(&user, "email = ?", loginBody.Email)
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

	// Build jwt
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	fmt.Println(tokenString, err)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to generate jwt token! Please try again.",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Session", tokenString, 3600*24*7, "", "", false, true)

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfuly",
		"data":    tokenString,
	})
}

func (a AuthController) Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	// Send response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Validated",
		"data":    user.(model.User).ID,
	})
}
