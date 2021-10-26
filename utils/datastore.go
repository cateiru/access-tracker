package utils

import "cloud.google.com/go/datastore"

func CreateKey(kind string, keys ...string) *datastore.Key {
	if len(keys) == 0 {
		return datastore.IncompleteKey(kind, nil)
	}

	return datastore.NameKey(kind, keys[0], nil)
}
