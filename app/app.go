package app

import (
	"github.com/eltrufas/jiritsu/db"
)

type App struct {
	DB     *db.Database
	Config *Config
}

func New(config *Config, db *db.Database) *App {
	return &App{db, config}
}
