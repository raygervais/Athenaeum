package controllers

import (
	"Golang/Athenaeum/src/backend/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
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

		AssertExpected(t, expected, response)
	})

	t.Run("Find Invalid Book ID", func(t *testing.T) {
		w, c := SetupContext(db)

		c.Params = []gin.Param{gin.Param{Key: "id", Value: "-2"}}

		FindBook(c)

		expected := "{\"error\":\"Record not found!\"}"

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, expected, w.Body.String())

	})

	t.Run("Create Valid Book", func(t *testing.T) {
		w, c := SetupContext(db)

		payload := models.Book{
			Title:  "Harry Potter and the Weird Sisters",
			Author: "J. K. Rowling",
		}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(payload)

		c.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes.Bytes())),
		}

		CreateBook(c)

		var response models.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Equal(t, nil, err)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, payload.Author, response.Author)
		assert.Equal(t, payload.Title, response.Title)
	})

	t.Run("Fail to Create Invalid Book", func(t *testing.T) {

	})
}
