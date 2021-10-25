package database

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Database struct {
	ctx    *context.Context
	client *datastore.Client
}

func New(ctx *context.Context, projectId string) (*Database, error) {
	client, err := datastore.NewClient(*ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &Database{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *Database) Close() {
	c.client.Close()
}

func (c *Database) Get(key *datastore.Key) (*datastore.Entity, error) {
	entity := new(datastore.Entity)

	if err := c.client.Get(*c.ctx, key, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (c *Database) Put(key *datastore.Key, entry datastore.Entity) error {
	if _, err := c.client.Put(*c.ctx, key, entry); err != nil {
		return err
	}

	return nil
}
