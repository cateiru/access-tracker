package utils

import "net/http"

// Get http query.
//
// Arguments
//	- r - request.
//	- key - query key.
//
// Returns
//	query string
func GetQuery(r *http.Request, key string) string {
	query := r.URL.Query().Get(key)

	return query
}
