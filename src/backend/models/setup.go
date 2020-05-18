package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels(dbTarget string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbTarget)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	return db
}
