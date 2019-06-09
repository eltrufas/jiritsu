package app

import (
	"context"

	"github.com/eltrufas/jiritsu/model"
)

func (app *App) CreateSource(ctx context.Context, source *model.Source) error {
	return app.DB.CreateSource(ctx, source)
}

func (app *App) ListSources(ctx context.Context, q model.PageQuery) (*model.Page, error) {
	return app.DB.ListSources(ctx, q)
}

func (app *App) GetSourceByID(ctx context.Context, id int64) (*model.Source, error) {
	return app.DB.GetSourceByID(ctx, id)
}

func (app *App) UpdateSource(ctx context.Context, source model.Source) error {
	return app.DB.UpdateSource(ctx, source)
}

func (app *App) DeleteSource(ctx context.Context, id int64) error {
	return app.DB.DeleteSource(ctx, id)
}
