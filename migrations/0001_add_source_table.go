package migrations

import (
	"github.com/jmoiron/sqlx"
	errors "golang.org/x/xerrors"
)

var createSourceTable = &Migration{
	Number: 2,
	Name:   "Create sources table",
	Up: func(tx *sqlx.Tx) error {
		const createTableSQL = `
			CREATE TABLE sources(
				id SERIAL PRIMARY KEY,
				name VARCHAR(1024) NOT NULL UNIQUE,
				path VARCHAR(4096) NOT NULL
			);
		`
		_, err := tx.Exec(createTableSQL)
		if err != nil {
			return errors.Errorf("Unable to create sources table: %w", err)
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, createSourceTable)
}
