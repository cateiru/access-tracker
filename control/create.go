package control

import (
	"context"

	"github.com/yuto51942/access-tracker/database"
	"github.com/yuto51942/access-tracker/types"
	"github.com/yuto51942/access-tracker/utils"
)

func Create(ctx *context.Context, redirectUrl string) ([]byte, error) {
	id, err := utils.CreateId()
	if err != nil {
		return nil, err
	}
	accessKey, err := utils.CreateId()
	if err != nil {
		return nil, err
	}

	if err := setDB(ctx, id, accessKey, redirectUrl); err != nil {
		return nil, err
	}

	return utils.ToJson(types.Created{TrackId: id[:8], AccessKey: accessKey, RedirectUrl: redirectUrl})
}

func setDB(ctx *context.Context, id string, accessKey string, redirectUrl string) error {
	db, err := database.New(ctx, database.ProjectID)
	if err != nil {
		return err
	}

	key := database.CreateKey("Tracking", id)
	historyKey := database.CreateKey("History", id)

	if err := db.Put(key, types.IdEntity{
		TrackId: id, AccessKey: accessKey, RedirectUrl: redirectUrl, History: historyKey}); err != nil {
		return err
	}

	return nil
}
