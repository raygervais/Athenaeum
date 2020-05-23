package models

// Collection is the internal DB structure which houses a collection of Books
type Collection struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Books       []Book `json:"books"`
}

// CreateCollectionInput is the JSON binding validation handler
type CreateCollectionInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Books       []Book `json:"books" gorm:"foreignkey:BookID`
}

// UpdateCollectionInput is the JSON binding validation handler
type UpdateCollectionInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Books       []Book `json:"books"`
}
