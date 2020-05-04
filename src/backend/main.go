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

    router.GET("/secret", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "May the 4th be with you!"
        })
    })

	return router
}
