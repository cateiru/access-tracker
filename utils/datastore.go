package utils

import "cloud.google.com/go/datastore"

func CreateKey(kind string, keys ...string) *datastore.Key {
	var key = new(datastore.Key)

	for _, keyId := range keys {
		key = datastore.NameKey(kind, keyId, key)
	}

	return key
}
