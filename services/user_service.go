package services

import (
	"errors"
	"io"
	"net/http"

	"golang-project/database"
	"golang-project/models"
	"golang-project/validators"

	"github.com/gin-gonic/gin"
)

func CreateUserService(c *gin.Context) (int, gin.H) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
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

	if err := database.DB.Create(&user).Error; err != nil {
		return http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		}
	}

	return http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Data created successfully",
		"data":    user,
	}
}

func GetUsersService(c *gin.Context) (int, gin.H) {
	var users []models.User
	if err := database.DB.
		Preload("Profile").
		Preload("Leave").
		Find(&users).Error; err != nil {
		return http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		}
	}

	return http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data retrieved successfully",
		"data":    users,
	}
}

func GetUserService(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.
		Preload("Profile").
		Preload("Leave").
		First(&user, id).Error; err != nil {
		return http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		}
	}

	return http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data retrieved successfully",
		"data":    user,
	}
}

func UpdateUserService(c *gin.Context) (int, gin.H) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.
		Preload("Profile").
		Preload("Leave").
		First(&user, id).Error; err != nil {
		return http.StatusNotFound, gin.H{"error": "User not found"}
	}

	var reqBody models.User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	// 3. Update basic user fields
	user.Name = reqBody.Name

	// 4. Update profile (if exists)
	if user.Profile.UserID != 0 {
		user.Profile.NationalID = reqBody.Profile.NationalID
		user.Profile.BirthPlace = reqBody.Profile.BirthPlace
		user.Profile.BirthDate = reqBody.Profile.BirthDate
		user.Profile.Address = reqBody.Profile.Address
	}

	// 5. Update leave balance (if exists)
	if user.Leave.UserID != 0 {
		user.Leave.AnnualLeave = reqBody.Leave.AnnualLeave
		user.Leave.SickLeave = reqBody.Leave.SickLeave
		user.Leave.UnpaidLeave = reqBody.Leave.UnpaidLeave
	}

	// Begin transaction
	tx := database.DB.Begin()

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, gin.H{"error": "Failed to update user"}
	}
	if err := tx.Save(&user.Profile).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, gin.H{"error": "Failed to update profile"}
	}
	if err := tx.Save(&user.Leave).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, gin.H{"error": "Failed to update leave balance"}
	}

	tx.Commit()

	// Fetch updated data
	var updatedUser models.User
	database.DB.Preload("Profile").Preload("Leave").First(&updatedUser, id)

	statusMsg := "Data updated successfully"

	return http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": statusMsg,
		"data":    updatedUser,
	}
}

func DeleteUserService(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return http.StatusNotFound, gin.H{"error": "User not found"}
	}

	database.DB.Delete(&user)

	return http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data deleted successfully",
		"data":    user,
	}
}
