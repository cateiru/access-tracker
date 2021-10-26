package control

import (
	"context"

	"github.com/yuto51942/access-tracker/database"
)

func Delete(ctx *context.Context, id string, accessKey string) error {
	dbOp, err := database.NewOperator(ctx, id, accessKey)
	if err != nil {
		return err
	}
	defer dbOp.Close()

	return dbOp.Delete()
}
