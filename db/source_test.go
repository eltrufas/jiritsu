package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/eltrufas/jiritsu/migrations"
	"github.com/eltrufas/jiritsu/model"
	"github.com/stretchr/testify/assert"
)

var db *Database

func TestMain(m *testing.M) {
	cfg := Config{
		ConnectionString: "host=127.0.0.1 user=postgres password=postgres sslmode=disable",
	}
	var err error
	db, err = New(&cfg)

	_, err = db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	if err != nil {
		log.Fatalf("Error clearing db: %s", err)
	}

	migrations.Migrate(db.DB)

	if err != nil {
		return
	}

	os.Exit(m.Run())
}

func TestCreateSource(t *testing.T) {
	source := model.Source{Name: "create_test", Path: "/test/path"}
	ctx := context.Background()
	err := db.CreateSource(ctx, &source)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}

	assert.NotEqual(t, 0, source.ID, "ID for source not set")
}

func TestGetSourceByID(t *testing.T) {
	sourceA := &model.Source{Name: "get_test", Path: "/test/path"}
	ctx := context.Background()
	err := db.CreateSource(ctx, sourceA)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}

	sourceB, err := db.GetSourceByID(ctx, sourceA.ID)
	if err != nil {
		t.Errorf("GetSourceByID returned an error for a just created row")
	}

	assert.Equal(t, sourceA, sourceB, "Source is different from created source")
}

func TestListSources(t *testing.T) {
	source := &model.Source{Name: "test_source_5", Path: "/test/path"}
	ctx := context.Background()
	err := db.CreateSource(ctx, source)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}
	source.Name = "test_source_6"
	err = db.CreateSource(ctx, source)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}
	source.Name = "test_source_7"
	err = db.CreateSource(ctx, source)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}
	source.Name = "test_source_8"
	err = db.CreateSource(ctx, source)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}

	pq := model.PageQuery{PageSize: 1, PageToken: 0}
	page, err := db.ListSources(ctx, pq)
	if err != nil {
		t.Errorf("ListSources failed")
	}
	assert.Equal(t, len(page.Data), 1, "Source is different from created source")

	pq.PageSize = 2
	page, err = db.ListSources(ctx, pq)
	if err != nil {
		t.Errorf("ListSources failed")
	}

	assert.Equal(t, len(page.Data), 2, "Source is different from created source")
}

func TestUpdateSource(t *testing.T) {
	sourceA := &model.Source{Name: "update_test", Path: "/test/path"}
	ctx := context.Background()
	err := db.CreateSource(ctx, sourceA)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}

	sourceA.Name = "update_test_2"
	sourceA.Path = "/other/path"
	err = db.UpdateSource(ctx, *sourceA)
	if err != nil {
		t.Errorf("UpdateSource failed: %s", err)
	}

	sourceB, err := db.GetSourceByID(ctx, sourceA.ID)
	if err != nil {
		t.Errorf("GetSourceByID returned an error for a just created row")
	}

	assert.Equal(t, sourceA, sourceB, "Source is different from created source")
}

func TestDeleteSource(t *testing.T) {
	source := &model.Source{Name: "delete_test", Path: "/test/path"}
	ctx := context.Background()
	err := db.CreateSource(ctx, source)
	if err != nil {
		t.Errorf("CreateSource failed: %s", err)
	}

	err = db.DeleteSource(ctx, source.ID)
	if err != nil {
		t.Errorf("DeleteSource failed: %s", err)
	}

	emptySource, err := db.GetSourceByID(ctx, source.ID)
	if err != nil {
		t.Errorf("GetSourceByID returned an error for a just created row")
	}

	assert.Nil(t, emptySource)
}
