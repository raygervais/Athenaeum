package controllers

import (
	"Golang/Athenaeum/src/backend/models"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBookCRUDFunctions(t *testing.T) {
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

		c.Params = []gin.Param{gin.Param{Key: "id", Value: "2"}}

		FindBook(c)

		expected := models.Book{
			ID:     2,
			Title:  "Harry Potter and The Chamber of Secrets",
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

		SetupRequestBody(c, payload)

		CreateBook(c)

		var response models.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Equal(t, nil, err)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, payload.Author, response.Author)
		assert.Equal(t, payload.Title, response.Title)
	})

	t.Run("Create Invalid Book With Missing Author", func(t *testing.T) {
		w, c := SetupContext(db)

		payload := models.Book{
			Title: "Harry Potter and the Sad Sisters",
		}

		SetupRequestBody(c, payload)

		CreateBook(c)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"Key: 'CreateBookInput.Author' Error:Field validation for 'Author' failed on the 'required' tag\"}", w.Body.String())
	})

	t.Run("Create Invalid Book With Missing Title", func(t *testing.T) {
		w, c := SetupContext(db)

		payload := models.Book{
			Author: "J. K. Rowling",
		}

		SetupRequestBody(c, payload)

		CreateBook(c)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"Key: 'CreateBookInput.Title' Error:Field validation for 'Title' failed on the 'required' tag\"}", w.Body.String())
	})

	t.Run("Create Invalid Book With Missing Fields", func(t *testing.T) {
		w, c := SetupContext(db)

		payload := models.Book{}

		SetupRequestBody(c, payload)

		CreateBook(c)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"Key: 'CreateBookInput.Title' Error:Field validation for 'Title' failed on the 'required' tag\\nKey: 'CreateBookInput.Author' Error:Field validation for 'Author' failed on the 'required' tag\"}", w.Body.String())
	})

	t.Run("Update Valid Book", func(t *testing.T) {
		w, c := SetupContext(db)

		payload := models.CreateBookInput{
			Title: "Hermione Granger and The Wibbly Wobbly Timey Wimey Escape",
		}

		SetupRequestBody(c, payload)
		c.Params = []gin.Param{gin.Param{Key: "id", Value: "3"}}

		UpdateBook(c)

		var response models.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, payload.Title, response.Title)
	})

	t.Run("Update Invalid Format", func(t *testing.T) {
		w, c := SetupContext(db)

		payload, _ := json.Marshal(map[int]string{
			2: "Harry Potter",
			3: "JK Rowling",
			4: "22",
		})

		SetupRequestBody(c, payload)
		c.Params = []gin.Param{gin.Param{Key: "id", Value: "3"}}

		UpdateBook(c)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"json: cannot unmarshal string into Go value of type models.UpdateBookInput\"}", w.Body.String())
	})

	t.Run("Update Invalid Book ID", func(t *testing.T) {
		w, c := SetupContext(db)

		payload := models.CreateBookInput{
			Title: "Hermione Granger and The Wibbly Wobbly Timey Wimey Escape",
		}

		SetupRequestBody(c, payload)
		c.Params = []gin.Param{gin.Param{Key: "id", Value: "-3"}}

		UpdateBook(c)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"Record not found!\"}", w.Body.String())
	})

	t.Run("Delete Valid Book ID", func(t *testing.T) {
		w, c := SetupContext(db)
		c.Params = []gin.Param{gin.Param{Key: "id", Value: "5"}}

		DeleteBook(c)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"data\":true}", w.Body.String())
	})

	t.Run("Delete Invalid Book ID", func(t *testing.T) {
		w, c := SetupContext(db)
		c.Params = []gin.Param{gin.Param{Key: "id", Value: "-2"}}

		DeleteBook(c)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"Record not found!\"}", w.Body.String())
	})
}
