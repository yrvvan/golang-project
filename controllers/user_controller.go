package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-project/services"
)

func GetUsers(c *gin.Context) {
	status, res := services.GetUsersService(c)
	c.JSON(status, res)
}

func GetUserByID(c *gin.Context) {
	status, res := services.GetUserService(c)
	c.JSON(status, res)
}

func CreateUser(c *gin.Context) {
	status, res := services.CreateUserService(c)
	c.JSON(status, res)
}

func UpdateUser(c *gin.Context) {
	status, res := services.UpdateUserService(c)
	c.JSON(status, res)
}

func DeleteUser(c *gin.Context) {
	status, res := services.DeleteUserService(c)
	c.JSON(status, res)
}
