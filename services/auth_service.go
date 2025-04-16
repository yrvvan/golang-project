package services

import (
	"errors"
	"io"
	"net/http"

	"golang-project/database"
	"golang-project/models"
	"golang-project/utils"
	"golang-project/validators"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) (int, gin.H) {
	var req struct {
		Email    string `binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}

	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.Is(err, io.EOF) {
			return http.StatusUnprocessableEntity, gin.H{
				"status":  http.StatusUnprocessableEntity,
				"message": "Invalid request payload",
			}
		}

		errors := validators.FormatValidationErrors(err)

		if len(errors) > 0 {
			return http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad request payload",
				"errors":  errors,
			}
		}
	}

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return http.StatusUnauthorized, gin.H{"error": "User not found"}
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return http.StatusUnauthorized, gin.H{"error": "Invalid credentials"}
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Failed to generate token"}
	}

	return http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Token generated successfully",
		"token":   token,
	}
}
