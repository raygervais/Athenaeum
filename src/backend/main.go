package main

import "github.com/gin-gonic/gin"

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

	return router
}
