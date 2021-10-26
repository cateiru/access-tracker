package types

import "cloud.google.com/go/datastore"

type IdEntity struct {
	TrackId     string
	AccessKey   string
	RedirectUrl string
	History     *datastore.Key
}
