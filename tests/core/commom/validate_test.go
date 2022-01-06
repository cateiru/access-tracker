package commom_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/access-tracker/core/common"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
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
		RedirectUrl: "hoge",
		Create:      time.Now(),
	}

	err = track.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetTrackByTrackID(ctx, db, trackId)
		require.NoError(t, err)

		return entity != nil
	}, "")

	err = common.ValidateKey(ctx, db, trackId, accessKey)
	require.NoError(t, err)

	err = common.ValidateKey(ctx, db, trackId, utils.CreateID(0))
	require.Error(t, err)

	err = common.ValidateKey(ctx, db, utils.CreateID(0), accessKey)
	require.Error(t, err)
}
