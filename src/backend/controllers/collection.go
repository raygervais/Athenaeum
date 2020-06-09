package controllers

import (
	"Golang/Athenaeum/src/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RetrieveCollectionByID is a helper function which returns a boolean based on success to find collection
func RetrieveCollectionByID(db *gorm.DB, c *gin.Context, book *models.Collection) bool {
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errRecordNotFound})
		return false
	}

	return true
}

// FindCollections called by GET /collections
// Get all collections
func FindCollections(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var collection models.Collection

	if !RetrieveCollectionByID(db, c, &collection) {
		return
	}

	c.JSON(http.StatusOK, collection)
}

// CreateCollection call by POST /collections
func CreateCollection(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.CreateCollectionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create collection
	collection := models.Collection{Title: input.Title, Description: input.Description, Books: input.Books}
	db.Create(&collection)

	c.JSON(http.StatusOK, collection)

}

// FindCollection call by GET /collections/:id
func FindCollection(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var collection models.Collection

	if !RetrieveCollectionByID(db, c, &collection) {
		return
	}

	c.JSON(http.StatusOK, collection)

}

// UpdateCollection call by PATCH /collections/:id
func UpdateCollection(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var collection models.Collection
	if !RetrieveCollectionByID(db, c, &collection) {
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&collection).Updates(input)

	c.JSON(http.StatusOK, collection)

}

// DeleteCollection call by /collections/:id
func DeleteCollection(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var collection models.Collection
	if !RetrieveCollectionByID(db, c, &collection) {
		return
	}

	db.Delete(&collection)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
