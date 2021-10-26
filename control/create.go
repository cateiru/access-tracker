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
	id = id[:10]
	accessKey, err := utils.CreateId()
	if err != nil {
		return nil, err
	}

	value := types.Created{TrackId: id, AccessKey: accessKey, RedirectUrl: redirectUrl}

	if err := setDB(ctx, value); err != nil {
		return nil, err
	}

	return utils.ToJson(value)
}

func setDB(ctx *context.Context, value types.Created) error {
	db, err := database.New(ctx, database.ProjectID)
	if err != nil {
		return err
	}

	key := database.CreateKey("Tracking", value.TrackId)
	historyKey := database.CreateKey("History", value.TrackId)

	if err := db.Put(key, types.IdEntity{
		TrackId:     value.TrackId,
		AccessKey:   value.AccessKey,
		RedirectUrl: value.RedirectUrl,
		History:     historyKey,
	}); err != nil {
		return err
	}

	return nil
}
