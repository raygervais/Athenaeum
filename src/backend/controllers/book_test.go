package controllers

import (
	"Golang/Athenaeum/src/backend/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func AssertExpected(t *testing.T, expected, received interface{}) {
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("Got %v, wanted %v", received, expected)
	}
}

func SetupContext(db *gorm.DB) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("db", db)

	return w, c
}

func TestCRUDFunctions(t *testing.T) {
	db := models.SetupModels("../test.db")

	gin.SetMode(gin.TestMode)

	t.Run("Query Test DB", func(t *testing.T) {
		w, c := SetupContext(db)

		FindBooks(c)

		AssertExpected(t, w.Code, http.StatusOK)
		AssertExpected(t, len(w.Body.String()) > 0, true)
	})

	t.Run("Find Valid Book", func(t *testing.T) {
		w, c := SetupContext(db)

		c.Params = []gin.Param{gin.Param{Key: "id", Value: "3"}}

		FindBook(c)

		expected := models.Book{
			ID:     3,
			Title:  "Harry Potter and The Prisoner of Azkaban",
			Author: "J. K. Rowling",
		}

		var response models.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
		AssertExpected(t, expected, response)
	})
}
