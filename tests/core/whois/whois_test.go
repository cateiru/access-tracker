package whois

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/access-tracker/core/whois"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestWhois(t *testing.T) {
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

	// --- 取得する

	histories, err := whois.WhoIs(ctx, trackId, accessKey, -1)
	require.NoError(t, err)
	require.Len(t, histories, 5)
}
