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

func TestGetBookById(t *testing.T) {
	collection := GenerateSampleBooks()

	t.Run("Get By Valid Id", func(t *testing.T) {
		expected := collection[0]
		received, _ := GetBook(0)

		AssertExpected(t, expected, received)
	})

	t.Run("Get by Invalid Id", func(t *testing.T) {
		_, err := GetBook(1000000)

		AssertExpected(t, errors.New("Id provided is greater than book list"), err)
	})

	t.Run("Get by Invalid Negative Id", func(t *testing.T) {
		_, err := GetBook(-1)

		AssertExpected(t, errors.New("Invalid id provided"), err)
	})

}
