package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/harrisonmalone/puppies-server/handlers"
	"github.com/harrisonmalone/puppies-server/models"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	r.GET("/puppies", handlers.PuppyRoutes{}.Index)
	r.POST("/puppies", handlers.PuppyRoutes{}.Create)
	r.GET("/puppies/:id", handlers.PuppyRoutes{}.Show)
	r.DELETE("/puppies/:id", handlers.PuppyRoutes{}.Destroy)
	r.POST("/users", handlers.UserRoutes{}.Create)
	return r
}

func main() {
	models.ConnectToDb()
	r := setupRouter()
	r.Run("localhost:8080")
}
