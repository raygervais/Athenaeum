package main

import (
	"Golang/Athenaeum/src/backend/models"
	"bytes"
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

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "Hello, World!",
	}

	router, db := SetupRouter("test.db")
	defer db.Close()
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

func TestBooksCRUD(t *testing.T) {
	dbTarget := "test.db"
	router, db := SetupRouter(dbTarget)
	db.DropTableIfExists(&models.Book{}, "books")
	db = models.SetupModels(dbTarget)

	defer db.Close()

	t.Run("Create Empty DB", func(t *testing.T) {
		w := performRequest(router, "GET", "/books/")

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Retrieve Nonexistent ID on Empty DB", func(t *testing.T) {

		w := performRequest(router, "GET", "/book/2")

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Populate DB with Harry Potter Set", func(t *testing.T) {
		books := []string{
			"Harry Potter and The Philosopher's Stone",
			"Harry Potter and The Chamber of Secrets",
			"Harry Potter and The Prisoner of Azkaban",
			"Harry Potter and The Goblet of Fire",
			"Harry Potter and The Order of The Phoenix",
			"Harry Potter and The Half-Blood Prince",
			"Harry Potter and The Deathly Hallows",
		}

		for _, book := range books {

			payload, _ := json.Marshal(models.CreateBookInput{
				Author: "J. K. Rowling",
				Title:  book,
			})

			req, err := http.NewRequest("POST", "/books", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, nil, err)
			assert.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Retrieve Existing ID on Populated DB", func(t *testing.T) {
		w := performRequest(router, "GET", "/books/2")

		expected := models.Book{
			Author: "J. K. Rowling",
			ID:     2,
			Title:  "Harry Potter and The Chamber of Secrets",
		}

		var response models.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, response)
	})

	t.Run("Attempt Updating Non-Existing ID", func(t *testing.T) {
		w := performRequest(router, "PATCH", "/books/-2")

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"error\":\"Record not found!\"}", w.Body.String())
	})

	t.Run("Updated Existing ID with Invalid Values", func(t *testing.T) {
		payload, _ := json.Marshal(map[int]string{
			2: "Harry Potter",
			3: "JK Rowling",
			4: "22",
		})

		req, err := http.NewRequest("PATCH", "/books/-2", bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Update Existing ID on Populated DB", func(t *testing.T) {
		payload, _ := json.Marshal(models.UpdateBookInput{
			Author: "J. K. Rowling",
			Title:  "Harry Potter and The Gaslight Anthem",
		})

		req, err := http.NewRequest("PATCH", "/books/6", bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Get Updated Book from Populated DB", func(t *testing.T) {
		expected := models.Book{
			Author: "J. K. Rowling",
			Title:  "Harry Potter and The Gaslight Anthem",
			ID:     6,
		}

		w := performRequest(router, "GET", "/books/6")

		var response models.Book
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, response)
	})

	t.Run("Delete Invalid Book from Populated DB", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/books/-1")
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Delete Without ID Book from Populated DB", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/books/")
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "404 page not found", w.Body.String())
	})

	t.Run("Delete valid Book from Populated DB", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/books/6")

		assert.Equal(t, "{\"data\":true}", w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
