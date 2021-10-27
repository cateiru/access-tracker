package control

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yuto51942/access-tracker/database"
	"github.com/yuto51942/access-tracker/utils"
)

func WhoIs(ctx *context.Context, id string, accessKey string) ([]byte, error) {
	dbOp, err := database.NewOperator(ctx, id, accessKey)
	if err != nil {
		return nil, err
	}
	defer dbOp.Close()

	history, err := dbOp.GetHistory()
	if err != nil {
		return nil, err
	}

	logrus.Info(*history)

	return utils.ToJson(*history)
}
