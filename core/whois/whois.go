package whois

import (
	"context"

	"github.com/yuto51942/access-tracker/core"
	"github.com/yuto51942/access-tracker/utils"
)

func WhoIs(ctx *context.Context, id string, accessKey string) ([]byte, error) {
	dbOp, err := core.NewOperator(ctx, id, accessKey)
	if err != nil {
		return nil, err
	}
	defer dbOp.Close()

	history, err := dbOp.GetHistory()
	if err != nil {
		return nil, err
	}

	return utils.ToJson(*history)
}
