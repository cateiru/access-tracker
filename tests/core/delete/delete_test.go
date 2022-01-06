package delete_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/access-tracker/core/delete"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	trackId := utils.CreateID(5)
	accessKey := utils.CreateID(0)

	track := models.Track{
		TrackId:     trackId,
		AccessKey:   accessKey,
		RedirectUrl: "https://example.com",
		Create:      time.Now(),
	}
	err = track.Add(ctx, db)
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		history := models.History{
			UniqueId:  utils.CreateID(0),
			TrackId:   trackId,
			Ip:        "192.0.2.0",
			UserAgent: "",
			Date:      time.Now(),
		}
		err = history.Add(ctx, db)
		require.NoError(t, err)
	}

	goretry.Retry(t, func() bool {
		historyCount, err := models.CountHistoriesByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return historyCount == 5
	}, "")

	// --- 削除する

	err = delete.Delete(ctx, trackId, accessKey)
	require.NoError(t, err)

	// --- 削除されたか確認する

	goretry.Retry(t, func() bool {
		histories, err := models.GetHistoriesByTrackID(ctx, db, trackId, -1)
		require.NoError(t, err)

		return len(histories) == 0
	}, "historyが全部削除されている")

	goretry.Retry(t, func() bool {
		track, err := models.GetTrackByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return track == nil
	}, "track infoが削除されている")
}
