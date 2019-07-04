package api

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (api *Api) ApiRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", api.CreateSource)

	router.Route("/:sourceID", func(r chi.Router) {
		r.Use(SourceCtx)

		r.Get("/", api.GetSource)
		r.Put("/", api.GetSource)
		r.Delete("/", api.GetSource)
	})

	return router
}

func (api *Api) SourceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		idstr := chi.URLParam(r, "sourceID")
		id, err := strconv.ParseInt(idstr, 10, 64)
		ctx = context.WithValue(ctx, "sourceID", id)
		if err != nil {
			http.Error(w, "Invalid source ID", 400)
		}

		source, err := api.App.GetSourceByID(ctx, id)
		if err != nil {
			http.Error(w, "Internal server error", 500)
		}

		if source == nil {
			http.Error(w, "Source not found", 404)
		}

		ctx = context.WithValue(ctx, "source", source)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (api *Api) CreateSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	source := model.Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		http.Error(w, "Invalid source", 400)
		return
	}

	err := api.App.CreateSource(ctx, &source)
	if err != nil {
		http.Error(w, "Could not create source", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).encode(source); err != nil {
		logrus.Error(err)
	}
}

func (api *Api) GetSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	source := ctx.Value("source").(*model.Source)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).encode(source); err != nil {
		logrus.Error(err)
	}
}

func (api *Api) UpdateSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	source := model.Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		http.Error(w, "Invalid source", 400)
		return
	}

	original := ctx.Value("source").(*model.Source)

	source.ID = original.ID

	err := api.App.UpdateSource(ctx, &source)
	if err != nil {
		http.Error(w, "Could not update source", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).encode(source); err != nil {
		logrus.Error(err)
	}
}

func (api *Api) DeleteSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	source := ctx.Value("source").(*model.Source)

	api.App.DeleteSource(ctx, source.ID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).encode(source); err != nil {
		logrus.Error(err)
	}
}
