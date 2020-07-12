package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raygervais/athenaeum/cmd/controllers"
	"github.com/raygervais/athenaeum/pkg/application"
)

func SetupRouter(app *application.Application) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/books", controllers.ListBooks(app))
	router.POST("/books", controllers.CreateBook(app))
	router.GET("/books/:id", controllers.FindBook(app))
	router.PATCH("/books/:id", controllers.UpdateBook(app))
	router.DELETE("/books/:id", controllers.DeleteBook(app))

	return router
}
