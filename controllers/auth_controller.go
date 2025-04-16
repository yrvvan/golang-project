package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-project/services"
)

func Login(c *gin.Context) {
	status, res := services.Login(c)
	c.JSON(status, res)
}
