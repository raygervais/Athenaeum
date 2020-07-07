package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func Get(databaseName string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		Client: db,
	}, nil
}

func (db *Database) InitializeTables(tables interface{}) error {
	return db.Client.AutoMigrate(tables)
}
