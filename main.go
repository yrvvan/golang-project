package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-project/database"
	"golang-project/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	database.Connect()

	r := gin.Default()

	routes.RegisterUserRoutes(r) // ⬅️ mount the routes here

	r.Run(":8089")
}
