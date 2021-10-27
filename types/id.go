package types

import (
	"time"

	"cloud.google.com/go/datastore"
)

type IdEntity struct {
	Key         *datastore.Key `datastore:"__key__"`
	TrackId     string         `datastore:"TrackId"`
	AccessKey   string         `datastore:"AccessKey,noindex"`
	RedirectUrl string         `datastore:"RedirectUrl,noindex"`
	Create      time.Time      `datastore:"Create"`
}
