package main

import (
	"Golang/Athenaeum/src/backend/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":3000")
}

// ErrorHandler provides basic error handling between the front-end and back-end
func ErrorHandler(c *gin.Context, err error, status int) {
	c.JSON(status, gin.H{"error": err.Error()})
}

// SetupRouter defines the routes and controllers which power our backend API
func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/book/", func(c *gin.Context) {
		c.JSON(http.StatusOK, controllers.GetBooks())
	})

	router.GET("/book/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			ErrorHandler(c, err, http.StatusInternalServerError)
			return
		}

		book, err := controllers.GetBook(id)

		if err != nil {
			ErrorHandler(c, err, http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, book)

	})

	router.DELETE("/book/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			ErrorHandler(c, err, http.StatusInternalServerError)
			return
		}

		books, err := controllers.DeleteBook(id)

		if err != nil {
			ErrorHandler(c, err, http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, books)
	})

	return router
}
