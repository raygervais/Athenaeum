package main

import (
	"Golang/Athenaeum/src/backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":3000")
}

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/book/", controllers.GetBooks)
	router.GET("/book/:id", controllers.GetBook)

	return router
}
