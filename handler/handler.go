package handler

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yuto51942/access-tracker/control"
	"github.com/yuto51942/access-tracker/utils"
)

// Tracking and redirect
func TrackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get url path.
	// Example: http://example.com/hoge -> hoge
	path := strings.FieldsFunc(r.URL.Path, func(r rune) bool {
		return r == '/'
	})

	if len(path) != 1 || len(path[0]) == 0 {
		logrus.Info(r.URL.Path)
		logrus.Info(path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := path[0]
	ip := r.RemoteAddr

	redirect, err := control.Tracking(&ctx, id, ip)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if utils.IsUrl(redirect) {
		http.Redirect(w, r, redirect, http.StatusMovedPermanently)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// Setting: Create url, reference access history and delete tracking url.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		WhoisHandler(w, r)
	case http.MethodPost:
		CreateHandler(w, r)
	case http.MethodDelete:
		DeleteHandler(w, r)
	}
}

// Create tracking url
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.ParseForm()
	redirectUrl := r.PostFormValue("redirect")

	bytes, err := control.Create(&ctx, redirectUrl)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// Delete tracking url
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.GetQuery(r, "id")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accessKey, err := utils.GetQuery(r, "key")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := control.Delete(&ctx, id, accessKey); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Reference access history.
func WhoisHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.GetQuery(r, "id")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accessKey, err := utils.GetQuery(r, "key")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	history, err := control.WhoIs(&ctx, id, accessKey)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if history != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(history)
	}
}
