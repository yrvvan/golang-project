package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/go-playground/validator/v10"
	"go-rose/validators"
	"go-rose/database"
	"go-rose/handlers"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	database.Connect()

	router := gin.Default()

	validate := validator.New()

	// Register custom validation for email uniqueness
	validate.RegisterValidation("emailnotexist", validators.EmailNotExist)

	// Register the custom validator with Gin context
	router.Use(func(c *gin.Context) {
		// Set custom validator to context
		c.Set("validator", validate)
		c.Next()
	})

	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	router.Run(":8089")
}
