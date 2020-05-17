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

const (
	errorInvalidID = "Invalid ID provided"
	errorGreaterID = "Provided ID is greater than book list size"
)

// TODO: Development Mock DB

// GenerateSampleBooks returns an array of predefined Books for testing usecases.
func GenerateSampleBooks() []Book {
	return []Book{
		{
			ID:      0,
			Name:    "Harry Potter and the Sorcerers Stone",
			Release: 1994,
		},
		{
			ID:      1,
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

	if err := BookIDResolver(id); err != nil {
		return Book{}, err
	}

	return books[id], nil
}

// DeleteBook handles removing the book with the corresponding Id from the datasource
func DeleteBook(id int) ([]Book, error) {
	books := GenerateSampleBooks()

	if err := BookIDResolver(id); err != nil {
		return []Book{}, err
	}

	return append(books[:id], books[id+1:]...), nil
}

// NOTE: This is a helper function while DB is not implemented

// BookIDResolver handles row length handling while we are using index as id (TEMP)
func BookIDResolver(id int) error {
	books := GenerateSampleBooks()

	if id < 0 {
		return errors.New(errorInvalidID)
	}

	if len(books)-1 < id {
		return errors.New(errorGreaterID)
	}

	return nil
}
