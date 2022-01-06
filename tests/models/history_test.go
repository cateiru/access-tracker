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

func TestHistory(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	trackId := utils.CreateID(5)

	// 10個入れる
	for i := 0; i < 10; i++ {
		history := models.History{
			UniqueId:  utils.CreateID(0),
			TrackId:   trackId,
			Ip:        "192.0.2.0",
			UserAgent: "",
			Date:      time.Now(),
		}

		err := history.Add(ctx, db)
		require.NoError(t, err)
	}

	goretry.Retry(t, func() bool {
		count, err := models.CountHistoriesByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return count == 10
	}, "10個入っている")

	histories, err := models.GetHistoriesByTrackID(ctx, db, trackId, -1)
	require.NoError(t, err)
	require.Len(t, histories, 10)

	historiesLimit, err := models.GetHistoriesByTrackID(ctx, db, trackId, 3)
	require.NoError(t, err)
	require.Len(t, historiesLimit, 3)

	err = models.DeleteHistoriesByTrackID(ctx, db, trackId)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		histories, err := models.GetHistoriesByTrackID(ctx, db, trackId, -1)
		require.NoError(t, err)

		return len(histories) == 0
	}, "削除されている")
}
