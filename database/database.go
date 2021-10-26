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

func (c *Database) Get(key *datastore.Key, entity interface{}) error {
	if err := c.client.Get(*c.ctx, key, entity); err != nil {
		return err
	}

	return nil
}

func (c *Database) Put(key *datastore.Key, entry interface{}) error {
	if _, err := c.client.Put(*c.ctx, key, entry); err != nil {
		return err
	}

	return nil
}

func (c *Database) Delete(key *datastore.Key) error {
	return c.client.Delete(*c.ctx, key)
}

func CreateKey(kind string, keys ...string) *datastore.Key {
	var key = new(datastore.Key)

	for _, keyId := range keys {
		key = datastore.NameKey(kind, keyId, key)
	}

	return key
}
