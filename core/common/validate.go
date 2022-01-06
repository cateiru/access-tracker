package common

import (
	"context"
	"errors"

	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/go-http-error/httperror/status"
)

// そのTrackIDのアクセス権があるかを調べます
func ValidateKey(ctx context.Context, db *database.Database, trackId string, key string) error {
	track, err := models.GetTrackByTrackID(ctx, db, trackId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if track == nil {
		return status.NewBadRequestError(errors.New("entity is not found")).Caller()
	}

	if track.AccessKey != key {
		return status.NewBadRequestError(errors.New("inaccessible due to lack of authority")).Caller()
	}

	return nil
}
