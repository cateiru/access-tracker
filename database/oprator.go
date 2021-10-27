package database

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/yuto51942/access-tracker/types"
	"github.com/yuto51942/access-tracker/utils"
)

type Operator struct {
	db        *Database
	id        string
	accessKey string
}

func NewOperator(ctx *context.Context, id string, accessKey string) (*Operator, error) {
	db, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return &Operator{
		db:        db,
		id:        id,
		accessKey: accessKey,
	}, nil
}

func (c *Operator) GetTracking() (*types.IdEntity, error) {
	key := utils.CreateKey("Tracking", c.id)
	entity := new(types.IdEntity)

	if err := c.db.Get(key, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (c *Operator) SetTracking(redirectUrl string) error {
	key := utils.CreateKey("Tracking", c.id)
	entity := types.IdEntity{
		TrackId:     c.id,
		AccessKey:   c.accessKey,
		RedirectUrl: redirectUrl,
		Create:      time.Now(),
	}

	if err := c.db.Put(key, &entity); err != nil {
		return err
	}

	return nil
}

func (c *Operator) GetHistory() (*[]types.History, error) {
	entity, err := c.GetTracking()
	if err != nil {
		return nil, err
	}

	if entity.AccessKey == c.accessKey {
		query := datastore.NewQuery(c.id)
		var posts []types.History

		_, err := c.db.GetAll(query, &posts)
		if err != nil {
			return nil, err
		}

		return &posts, err
	}

	return nil, errors.New("access key is different")
}

func (c *Operator) SetHistory(ip string) error {
	uniqueId, err := utils.CreateId()
	if err != nil {
		return err
	}

	historyKey := utils.CreateKey(c.id, uniqueId)
	entity := types.History{
		Ip:       ip,
		UniqueId: uniqueId,
		Date:     time.Now(),
	}

	if err := c.db.Put(historyKey, &entity); err != nil {
		return err
	}

	return nil
}

func (c *Operator) Delete() error {
	histories, err := c.GetHistory()
	if err != nil {
		return err
	}

	keys := []*datastore.Key{}

	for _, history := range *histories {
		keys = append(keys, utils.CreateKey(c.id, history.UniqueId))
	}

	return c.db.Delete(keys)
}

func (c *Operator) Close() {
	c.db.Close()
}
