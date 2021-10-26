package control

import (
	"context"
	"time"

	"github.com/yuto51942/access-tracker/database"
	"github.com/yuto51942/access-tracker/types"
	"github.com/yuto51942/access-tracker/utils"
)

func Tracking(ctx *context.Context, id string, ip string) (string, error) {
	db, err := database.New(ctx, database.ProjectID)
	if err != nil {
		return "", err
	}

	key := database.CreateKey("Tracking", id)
	entity := new(types.IdEntity)

	if err := db.Get(key, entity); err != nil {
		return "", err
	}

	uniqueId, err := utils.CreateId()
	if err != nil {
		return "", err
	}

	historyKey := database.CreateKey("History", id, uniqueId)

	if err := db.Put(historyKey, types.History{
		Ip:       ip,
		UniqueId: uniqueId,
		Date:     time.Now(),
	}); err != nil {
		return "", err
	}

	return entity.RedirectUrl, nil
}
