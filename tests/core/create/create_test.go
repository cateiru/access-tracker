package create_test

import (
	"context"
	"testing"

	"github.com/cateiru/access-tracker/core/create"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	redirectUrl := "https://example.com"

	resp, err := create.Create(ctx, redirectUrl)
	require.NoError(t, err)

	require.Equal(t, resp.RedirectUrl, redirectUrl)
	require.NotEmpty(t, resp.AccessKey)
	require.NotEmpty(t, resp.TrackId)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	goretry.Retry(t, func() bool {
		track, err := models.GetTrackByTrackID(ctx, db, resp.TrackId)
		require.NoError(t, err)

		return track != nil && track.RedirectUrl == redirectUrl
	}, "DBに格納されている")
}
