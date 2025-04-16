package routes

import (
	"golang-project/controllers"
	"golang-project/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	// 🔒 Group: protected user routes
	users := r.Group("/users")
	users.Use(middlewares.AuthMiddleware()) // Apply JWT middleware

	users.GET("/", controllers.GetUsers)
	users.GET("/:id", controllers.GetUserByID)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
	users.POST("/", controllers.CreateUser)

	// 🔓 Public route
	r.POST("/login", controllers.Login)
}
