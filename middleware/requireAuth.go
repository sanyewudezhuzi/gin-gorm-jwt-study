package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/initializers"
	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(ctx *gin.Context) {
	// Get the cookie off req
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		ctx.Set("user", user)

		// Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
