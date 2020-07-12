package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raygervais/athenaeum/cmd/models"
	"github.com/raygervais/athenaeum/pkg/application"
	"gorm.io/gorm"
)

// RetrieveBookByID is a helper function which returns a boolean based on success to find book
func RetrieveBookByID(db *gorm.DB, c *gin.Context, book *models.Book) bool {
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return false
	}

	return true
}

func ListBooks(app *application.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		app.Database.Client.Find(&books)
		c.JSON(http.StatusOK, books)
	}
}

func CreateBook(app *application.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.CreateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book := models.Book{Title: input.Title, Author: input.Author}

		app.Database.Client.Create(&book)
		c.JSON(http.StatusOK, book)
	}
}

func FindBook(app *application.Application) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get model if exist
		var book models.Book

		if !RetrieveBookByID(app.Database.Client, c, &book) {
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

// UpdateBook called by PATCH /books/:id
func UpdateBook(app *application.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		// Get model if exist
		var book models.Book
		if !RetrieveBookByID(app.Database.Client, c, &book) {
			return
		}

		// Validate input
		var input models.UpdateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Model(&book).Updates(input)

		c.JSON(http.StatusOK, book)
	}
}

// DeleteBook called by DELETE /books/:id
// Delete a book
func DeleteBook(app *application.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		// Get model if exist
		var book models.Book
		if !RetrieveBookByID(app.Database.Client, c, &book) {
			return
		}

		db.Delete(&book)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
