package track_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/access-tracker/core/track"
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
	redirectUrl := "https://example.com"

	trackInfo := models.Track{
		TrackId:     trackId,
		AccessKey:   utils.CreateID(0),
		RedirectUrl: redirectUrl,
		Create:      time.Now(),
	}
	err = trackInfo.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		trackInfo, err := models.GetTrackByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return trackInfo != nil
	}, "")

	url, err := track.Tracking(ctx, trackId, "192.0.2.0", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")
	require.NoError(t, err)
	require.Equal(t, url, redirectUrl)

	goretry.Retry(t, func() bool {
		histories, err := models.GetHistoriesByTrackID(ctx, db, trackId, -1)
		require.NoError(t, err)

		return len(histories) == 1 && histories[0].TrackId == trackId
	}, "履歴が保存されている")
}

func TestTrackBot(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	trackId := utils.CreateID(5)
	redirectUrl := "https://example.com"

	trackInfo := models.Track{
		TrackId:     trackId,
		AccessKey:   utils.CreateID(0),
		RedirectUrl: redirectUrl,
		Create:      time.Now(),
	}
	err = trackInfo.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		trackInfo, err := models.GetTrackByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return trackInfo != nil
	}, "")

	url, err := track.Tracking(ctx, trackId, "192.0.2.0", "Googlebot")
	require.NoError(t, err)
	require.Equal(t, url, redirectUrl)

	goretry.Retry(t, func() bool {
		histories, err := models.GetHistoriesByTrackID(ctx, db, trackId, -1)
		require.NoError(t, err)

		return len(histories) == 0
	}, "Botは履歴に保存しない")
}
