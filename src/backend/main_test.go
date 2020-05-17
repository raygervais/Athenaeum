package main

import (
	"Golang/Athenaeum/src/backend/controllers"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Referencing https://medium.com/@craigchilds94/testing-gin-json-responses-1f258ce3b0b1
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type errResponse struct {
	Error   string `json:"error"`
	Context string `json:"context"`
}

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "Hello, World!",
	}

	router := SetupRouter()
	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Make some assertions on the correctness of the response.
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)
}

func TestBookCRUD(t *testing.T) {
	router := SetupRouter()

	// Redundant using same source, but this is while we don't have Mock DB connection
	collection := controllers.GenerateSampleBooks()

	t.Run("GET All", func(t *testing.T) {
		w := performRequest(router, "GET", "/book/")

		// Assert we encoded correctly,
		assert.Equal(t, http.StatusOK, w.Code)

		// Convert the JSON response to struct
		var response []controllers.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, collection, response)
	})

	t.Run("GET By Valid ID", func(t *testing.T) {
		w := performRequest(router, "GET", "/book/0")
		assert.Equal(t, http.StatusOK, w.Code)

		var response controllers.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, collection[0], response)
	})

	t.Run("GET by Invalid ID", func(t *testing.T) {
		w := performRequest(router, "GET", "/book/3")
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response errResponse
		var expectedMessage = errResponse{
			Error: "Provided ID is greater than book list size",
		}

		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, expectedMessage, response)
	})

	t.Run("GET by Invalid Negative ID", func(t *testing.T) {
		w := performRequest(router, "GET", "/book/-2")
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response errResponse
		var expectedMessage = errResponse{
			Error: "Invalid ID provided",
		}

		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, expectedMessage, response)
	})

	t.Run("DELETE by valid ID", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/book/0")
		assert.Equal(t, http.StatusOK, w.Code)

		var response []controllers.Book
		var expected = []controllers.Book{
			{
				ID:      1,
				Name:    "Harry Potter and the Chamber of Secrets",
				Release: 1998,
			},
		}

		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, expected, response)
	})

	t.Run("DELETE by Invalid Negative ID", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/book/-1")
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response errResponse
		var expected = errResponse{
			Error: "Invalid ID provided",
		}

		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, expected, response)
	})

	t.Run("DELETE by invalid ID", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/book/3")
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response errResponse
		var expected = errResponse{
			Error: "Provided ID is greater than book list size",
		}

		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, expected, response)
	})
}
