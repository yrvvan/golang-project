package handlers

import (
	"golang-project/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	status, response := services.GetUsersService(c)
	c.JSON(status, response)
}

func CreateUser(c *gin.Context){
	status, response := services.CreateUserService(c)
	c.JSON(status, response)
}

func GetUser(c *gin.Context) {
	status, response := services.GetUserService(c)
	c.JSON(status, response)
}

func UpdateUser(c *gin.Context) {
	status, response := services.UpdateUserService(c)
	c.JSON(status, response)
}

func DeleteUser(c *gin.Context) {
	status, response := services.DeleteUserService(c)
	c.JSON(status, response)
}
