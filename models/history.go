package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/access-tracker/database"
)

func GetHistoriesByTrackID(ctx context.Context, db *database.Database, trackId string) ([]History, error) {
	query := datastore.NewQuery("History").Filter("trackId =", trackId)
	var entities []History

	if _, err := db.GetAll(ctx, query, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

func DeleteHistoriesByTrackID(ctx context.Context, db *database.Database, trackId string) error {
	query := datastore.NewQuery("History").Filter("trackId =", trackId)
	var entities []History

	keys, err := db.GetAll(ctx, query, &entities)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func (c *History) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("History", c.UniqueId)
	return db.Put(ctx, key, c)
}
