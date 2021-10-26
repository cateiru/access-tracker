package control

import (
	"context"
	"errors"

	"github.com/yuto51942/access-tracker/database"
	"github.com/yuto51942/access-tracker/types"
)

func Delete(ctx *context.Context, id string, accessKey string) error {
	db, err := database.New(ctx, database.ProjectID)
	if err != nil {
		return err
	}

	key := database.CreateKey("Tracking", id)
	entity := new(types.IdEntity)

	if err := db.Get(key, entity); err != nil {
		return err
	}

	if entity.AccessKey == accessKey {
		// delete history
		if err := db.Delete(entity.History); err != nil {
			return err
		}

		// delete primary
		if err := db.Delete(key); err != nil {
			return err
		}
		return nil
	}

	return errors.New("access key is different")
}
