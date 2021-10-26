package control

import (
	"context"

	"github.com/yuto51942/access-tracker/database"
)

func Tracking(ctx *context.Context, id string, ip string) (string, error) {
	// note: accessKey is not used.
	dbOp, err := database.NewOperator(ctx, id, "")
	defer dbOp.Close()
	if err != nil {
		return "", err
	}

	// check exist
	entity, err := dbOp.GetTracking()
	if err != nil {
		return "", err
	}

	if err := dbOp.SetHistory(ip); err != nil {
		return "", err
	}

	return entity.RedirectUrl, nil
}
