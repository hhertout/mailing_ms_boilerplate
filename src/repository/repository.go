package repository

import (
	"database/sql"
	"mailer_ms/cmd/database"
)

type Repository struct {
	dbPool *sql.DB
}

func NewRepository() (*Repository, error) {
	dbPool, err := database.Connect()
	if err != nil {
		return nil, err
	}

	return &Repository{
		dbPool,
	}, nil
}
