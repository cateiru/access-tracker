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
	dbOp, err := database.NewOperator(ctx, value.TrackId, value.AccessKey)
	defer dbOp.Close()
	if err != nil {
		return err
	}

	return dbOp.SetTracking(value.RedirectUrl)
}
