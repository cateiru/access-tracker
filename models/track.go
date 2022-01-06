package models

import (
	"context"

	"github.com/cateiru/access-tracker/database"
)

func GetTrackByTrackID(ctx context.Context, db *database.Database, trackId string) (*Track, error) {
	key := database.CreateNameKey("Track", trackId)

	var entity Track

	empty, err := db.Get(ctx, key, &entity)
	if err != nil {
		return nil, err
	}
	if empty {
		return nil, nil
	}

	return &entity, nil
}

func DeleteTrackByTrackID(ctx context.Context, db *database.Database, trackId string) error {
	key := database.CreateNameKey("Track", trackId)

	return db.Delete(ctx, key)
}

func (c *Track) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("Track", c.TrackId)
	return db.Put(ctx, key, c)
}
