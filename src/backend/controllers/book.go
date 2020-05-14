package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Release int    `json:"release"`
}

// TODO: Development Mock DB
func GenerateSampleBooks() []Book {
	return []Book{
		{
			ID:      1,
			Name:    "Harry Potter and the Sorcerers Stone",
			Release: 1994,
		},
		{
			ID:      2,
			Name:    "Harry Potter and the Chamber of Secrets",
			Release: 1998,
		},
	}
}

/**
 * TODO: Seperate Controller logic from API Logic in main.go
 * Only controller logic should be here, such as handling structs etc.
 */
// API
func GetBooks(c *gin.Context) {
	books := GenerateSampleBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(fmt.Errorf("Error converting id to int"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	books := GenerateSampleBooks()

	if len(books)-1 < id {
		errorString := "Id provided is greater than book list"
		c.Error(errors.New(errorString))
		c.JSON(http.StatusNotFound, gin.H{"error": errorString})
		return
	}

	c.JSON(http.StatusOK, books[id])
}
