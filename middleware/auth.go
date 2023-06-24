package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/BitInByte/web-app-example/core"
	"github.com/BitInByte/web-app-example/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthGuard(ctx *gin.Context) {
	// Get session cookie
	tokenString, err := ctx.Cookie("Session")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	// Decode/validate
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find user with token sub
		var user model.User
		result := core.DB.First(&user, "email = ?", claims["sub"])
		if result.RowsAffected == 0 || result.Error != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		ctx.Set("user", user)

		// fmt.Println("In middleware")
		// Continue
		ctx.Next()
	} else {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
