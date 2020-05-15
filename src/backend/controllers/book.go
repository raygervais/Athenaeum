/*
Package controllers/books implements Book object CRUD patterns.
*/

package controllers

import (
	"errors"
)

// Book struct for Milestone 2, to be replaced with far grander models
type Book struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Release int    `json:"release"`
}

// TODO: Development Mock DB

// GenerateSampleBooks returns an array of predefined Books for testing usecases.
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

// GetBooks handles retrieving all Books from the DB and returning them to the server context
func GetBooks() []Book {
	return GenerateSampleBooks()
}

// GetBook handles retrieving only the book with the corresponding id that matches the provided parameter
func GetBook(id int) (Book, error) {
	books := GenerateSampleBooks()

	if id < 0 {
		return Book{}, errors.New("Invalid id provided")
	}

	if len(books)-1 < id {
		return Book{}, errors.New("Id provided is greater than book list")
	}

	return books[id], nil
}
