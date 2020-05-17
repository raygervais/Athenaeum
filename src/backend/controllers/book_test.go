package controllers

import (
	"errors"
	"reflect"
	"testing"
)

func AssertExpected(t *testing.T, expected, received interface{}) {
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("Got %v, wanted %v", received, expected)
	}
}

func TestGetBooks(t *testing.T) {
	expected := GenerateSampleBooks()
	received := GetBooks()

	AssertExpected(t, expected, received)
}

func TestGetBookByID(t *testing.T) {
	collection := GenerateSampleBooks()

	t.Run("Get By Valid ID", func(t *testing.T) {
		expected := collection[0]
		received, _ := GetBook(0)

		AssertExpected(t, expected, received)
	})

	t.Run("Get by Invalid ID", func(t *testing.T) {
		_, err := GetBook(1000000)

		AssertExpected(t, errors.New("Provided ID is greater than book list size"), err)
	})

	t.Run("Get by Invalid Negative ID", func(t *testing.T) {
		_, err := GetBook(-1)

		AssertExpected(t, errors.New("Invalid ID provided"), err)
	})
}

func TestDeleteBookByID(t *testing.T) {
	var expected = []Book{
		{
			ID:      1,
			Name:    "Harry Potter and the Chamber of Secrets",
			Release: 1998,
		},
	}

	t.Run("Delete By Valid ID", func(t *testing.T) {
		booklist, err := DeleteBook(0)

		AssertExpected(t, nil, err)
		AssertExpected(t, expected, booklist)
	})

	t.Run("Delete by Invalid ID", func(t *testing.T) {
		_, err := DeleteBook(3)

		AssertExpected(t, errors.New("Provided ID is greater than book list size"), err)
	})

	t.Run("Get by Invalid Negative ID", func(t *testing.T) {
		_, err := DeleteBook(-1)

		AssertExpected(t, errors.New("Invalid ID provided"), err)
	})
}
