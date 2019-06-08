package migrations

import (
	"database/sql"
	"sort"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	errors "golang.org/x/xerrors"
)

type Migration struct {
	Number uint
	Name   string

	Up func(tx *sqlx.Tx) error
}

func CreateMigrationsTable(db *sqlx.DB) error {
	const createTableSQL = `
		CREATE TABLE IF NOT EXISTS migrations(
			number SERIAL PRIMARY KEY,
			name VARCHAR(1024) NOT NULL
		);
	`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return errors.Errorf("Unable to create migrations table: %w", err)
	}
	return nil
}

func GetExistingMigrations(db *sqlx.DB) ([]Migration, error) {
	ms := []Migration{}
	err := db.Select(&ms, `
		SELECT * FROM migrations
		ORDER BY number DESC
	`)

	return ms, err
}

func MigrationExists(db *sqlx.DB, m Migration) (bool, error) {
	err := db.Get(&m, "SELECT * FROM migrations WHERE number=$1", m.Number)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}

func WriteMigration(tx *sqlx.Tx, m Migration) error {
	_, err := tx.Exec("INSERT INTO migrations(number, name) VALUES ($1, $2)", m.Number, m.Name)
	return err
}

func Migrate(db *sqlx.DB) error {
	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Number < Migrations[j].Number
	})

	logrus.Debug("ensuring migrations table is present")
	if err := CreateMigrationsTable(db); err != nil {
		return errors.Errorf("unable to create migrations table: %w", err)
	}

	for _, m := range Migrations {
		exists, err := MigrationExists(db, *m)
		if err != nil {
			return errors.Errorf("Unable to retrieve migration: %w", err)
		}

		if exists {
			continue
		}

		logrus.Infof("Applying %d", m.Number)
		tx, err := db.Beginx()
		if err != nil {
			return errors.Errorf("Could not create db transaction: %w", err)
		}

		if err := m.Up(tx); err != nil {
			tx.Rollback()
			return errors.Errorf("Unable to apply migration. Rolling back: %w", err)
		}

		if err := WriteMigration(tx, *m); err != nil {
			tx.Rollback()
			return errors.Errorf("Unable set migration in db. Rolling back: %w", err)
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			return errors.Errorf("Unable to apply migration. Rolling back: %w", err)
		}
	}

	return nil
}

var Migrations []*Migration
