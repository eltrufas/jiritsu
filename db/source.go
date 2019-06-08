package db

import (
	"context"
	"database/sql"

	"github.com/eltrufas/jiritsu/model"
	errors "golang.org/x/xerrors"
)

func (db *Database) CreateSource(ctx context.Context, source *model.Source) error {
	err := db.QueryRowContext(
		ctx,
		"INSERT INTO sources(name, path) VALUES ($1, $2) RETURNING id",
		source.Name,
		source.Path,
	).Scan(&source.ID)
	if err != nil {
		return errors.Errorf("Unable to create source: %w", err)
	}
	return nil
}

func (db *Database) ListSources(ctx context.Context, q model.PageQuery) (*model.Page, error) {
	sources := []model.Source{}
	err := db.SelectContext(
		ctx,
		&sources,
		"SELECT * FROM sources WHERE id > $1 LIMIT $2",
		q.PageToken,
		q.PageSize,
	)
	if err != nil {
		return nil, errors.Errorf("Unable to retrieve sources: %w", err)
	}
	page := &model.Page{Data: make([]model.Model, len(sources))}
	for i, s := range sources {
		page.Data[i] = s
	}
	return page, err
}

func (db *Database) GetSourceByID(ctx context.Context, id int64) (*model.Source, error) {
	source := model.Source{}
	err := db.GetContext(
		ctx,
		&source,
		"SELECT * FROM sources WHERE id=$1 limit 1",
		id,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Errorf("Unable to retrieve source from DB: %w", err)
	}
	return &source, nil
}
