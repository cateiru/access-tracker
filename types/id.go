package types

import (
	"time"

	"cloud.google.com/go/datastore"
)

type IdEntity struct {
	TrackId     string         `datastore:"TrackId"`
	AccessKey   string         `datastore:"AccessKey,noindex"`
	RedirectUrl string         `datastore:"RedirectUrl,noindex"`
	History     *datastore.Key `datastore:"History"`
	Create      time.Time      `datastore:"Create"`
}
