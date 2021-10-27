package database

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Database struct {
	ctx    *context.Context
	client *datastore.Client
}

func New(ctx *context.Context) (*Database, error) {
	client, err := datastore.NewClient(*ctx, "")
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
	return c.client.Get(*c.ctx, key, entity)
}

func (c *Database) GetAll(query *datastore.Query, entities interface{}) ([]*datastore.Key, error) {
	return c.client.GetAll(*c.ctx, query, entities)
}

func (c *Database) Put(key *datastore.Key, entry interface{}) error {
	if _, err := c.client.Put(*c.ctx, key, entry); err != nil {
		return err
	}

	return nil
}

func (c *Database) DeleteMulti(key []*datastore.Key) error {
	return c.client.DeleteMulti(*c.ctx, key)
}

func (c *Database) Delete(key *datastore.Key) error {
	return c.client.Delete(*c.ctx, key)
}
