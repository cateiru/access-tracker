package control

import (
	"context"

	"github.com/yuto51942/access-tracker/database"
)

func Delete(ctx *context.Context, id string, accessKey string) error {
	dbOp, err := database.NewOperator(ctx, id, accessKey)
	defer dbOp.Close()
	if err != nil {
		return err
	}

	return dbOp.Delete()
}
