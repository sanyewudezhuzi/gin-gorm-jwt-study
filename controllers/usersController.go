package controllers

import (
	"net/http"

	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/initializers"
	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}
	if ctx.ShouldBindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{})
}
