package handler

import (
	"net/http"

	"github.com/yuto51942/access-tracker/control"
	"github.com/yuto51942/access-tracker/utils"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetQuery(r, "id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
