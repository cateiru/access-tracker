package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestTrack(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	trackId := utils.CreateID(5)

	entity := models.Track{
		TrackId:     trackId,
		AccessKey:   utils.CreateID(0),
		RedirectUrl: "https://example.com",
		Create:      time.Now(),
	}
	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		track, err := models.GetTrackByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return track != nil && track.TrackId == trackId
	}, "格納されている")

	err = models.DeleteTrackByTrackID(ctx, db, trackId)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		track, err := models.GetTrackByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return track == nil
	}, "削除されている")
}
