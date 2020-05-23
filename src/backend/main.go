package main

import (
	"Golang/Athenaeum/src/backend/controllers"
	"Golang/Athenaeum/src/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	// Configure the back-end router
	router, db := SetupRouter("local.db")
	defer db.Close()

	router.Run(":3000")
}

// SetupRouter defines the routes and controllers which power our backend API
func SetupRouter(dbTarget string) (*gin.Engine, *gorm.DB) {
	//
	router := gin.Default()

	// Create SQLite Database if it doesn't exist
	db := models.SetupModels(dbTarget)

	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/books/", controllers.FindBooks)
	router.POST("/books/", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.GET("/collections/", controllers.FindCollections)
	router.POST("/collections/", controllers.CreateCollection)
	router.GET("/collections/:id", controllers.FindCollection)
	router.PATCH("/collections/:id", controllers.UpdateCollection)
	router.DELETE("/collections/:id", controllers.DeleteCollection)
	return router, db
}
