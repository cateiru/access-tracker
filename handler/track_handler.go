package handler

import (
	"net/http"

	"github.com/cateiru/access-tracker/core/track"
	"github.com/cateiru/access-tracker/utils/net"
)

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	if err := track.TrackHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
