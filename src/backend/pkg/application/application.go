package application

import (
	"github.com/raygervais/athenaeum/pkg/config"
	"github.com/raygervais/athenaeum/pkg/db"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	Database *db.Database
	Config   *config.Config
}

// Get captures env vars, establishes DB connection and keeps/returns
// reference to both
func Get() (*Application, error) {

	db, err := db.Get(config.Get().GetDatabaseName("live"))
	if err != nil {
		return nil, err
	}

	return &Application{
		Database: db,
		Config:   config.Get(),
	}, nil
}
