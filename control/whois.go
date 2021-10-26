package control

import (
	"context"

	"github.com/yuto51942/access-tracker/database"
	"github.com/yuto51942/access-tracker/utils"
)

func WhoIs(ctx *context.Context, id string, accessKey string) ([]byte, error) {
	dbOp, err := database.NewOperator(ctx, id, accessKey)
	if err != nil {
		return nil, err
	}

	history, err := dbOp.GetHistory()
	if err != nil {
		return nil, err
	}

	return utils.ToJson(history)
}
