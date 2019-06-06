package db

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func New(config *Config) (*Database, error) {
	db, err := sqlx.Open("postgres", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
