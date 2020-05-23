package controllers

import (
	"Golang/Athenaeum/src/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RetrieveCollectionByID(db *gorm.DB, c *gin.Context, book *models.Book) bool {
	return true
}

func FindCollections(c *gin.Context) {
}

func CreateCollection(c *gin.Context) {

}

func FindCollection(c *gin.Context) {

}

func UpdateCollection(c *gin.Context) {

}

func DeleteCollection(c *gin.Context) {

}
