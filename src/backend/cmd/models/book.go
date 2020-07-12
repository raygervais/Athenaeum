package models

// Book is the internal DB structure
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// CreateBookInput is the JSON binding validation handler
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput is the JSON binding validation handler
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
