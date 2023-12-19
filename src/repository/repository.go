package repository

import (
	"database/sql"
	"mailer_ms/cmd/database"
)

type Repository struct {
	dbPool *sql.DB
}

func NewRepository(customSource *sql.DB) (*Repository, error) {
	if customSource != nil {
		return &Repository{
			customSource,
		}, nil
	} else {
		dbPool, err := database.Connect()
		if err != nil {
			return nil, err
		}

		return &Repository{
			dbPool,
		}, nil
	}
}
