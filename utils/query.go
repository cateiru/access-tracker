package utils

import (
	"errors"
	"net/http"
)

// Get http query.
//
// Arguments
//	- r - request.
//	- key - query key.
//
// Returns
//	query string
func GetQuery(r *http.Request, key string) (string, error) {
	query := r.URL.Query().Get(key)

	if len(query) == 0 {
		return "", errors.New("query is empty")
	}

	return query, nil
}
