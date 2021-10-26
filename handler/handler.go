package handler

import (
	"net/http"
	"strings"

	"github.com/yuto51942/access-tracker/control"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// get url path.
	// Example: http://example.com/hoge -> hoge
	path := strings.FieldsFunc(r.URL.Path, func(r rune) bool {
		return r == '/'
	})

	if len(path) != 1 || len(path[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := path[0]

	w.Write([]byte(id))
}

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := control.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func WhoisHandler(w http.ResponseWriter, r *http.Request) {

}
