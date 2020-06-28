package controllers

import (
	"Golang/Athenaeum/src/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	errRecordNotFound = "Record not found!"
)

// RetrieveBookByID is a helper function which returns a boolean based on success to find book
func RetrieveBookByID(db *gorm.DB, c *gin.Context, book *models.Book) bool {
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errRecordNotFound})
		return false
	}

	return true
}

// FindBooks called by GET /books
// Get all books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, books)
}

// CreateBook called by POST /books
// Create new book
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, book)
}

// FindBook called by GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var book models.Book

	if !RetrieveBookByID(db, c, &book) {
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook called by PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var book models.Book
	if !RetrieveBookByID(db, c, &book) {
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

// DeleteBook called by DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var book models.Book
	if !RetrieveBookByID(db, c, &book) {
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
