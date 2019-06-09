package api

import (
	"github.com/eltrufas/jiritsu/app"
	"github.com/go-chi/chi"
)

type API struct {
	App    *app.App
	Config *Config
}

func (api *API) Routes() chi.Router {
	router := chi.NewRouter()

	return router
}

func New(config *Config, app *app.App) *API {
	return &API{
		App:    app,
		Config: config,
	}
}
