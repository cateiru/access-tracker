package routes

import (
	"net/http"

	"github.com/cateiru/access-tracker/handler"
)

func Defs(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.TrackHandler)
	mux.HandleFunc("/u", handler.UserHandler)
}
