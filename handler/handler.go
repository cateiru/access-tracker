package handler

import (
	"net/http"

	"github.com/cateiru/access-tracker/core/create"
	"github.com/cateiru/access-tracker/core/delete"
	"github.com/cateiru/access-tracker/core/whois"
	"github.com/cateiru/access-tracker/utils/net"
)

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

	w.Header().Set("Cache-Control", "no-store")
}

// tracking urlを作成する
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := create.CreateHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// tracking urlを削除する
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := delete.DeleteHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// tracking urlのアクセスログを取得する
func WhoisHandler(w http.ResponseWriter, r *http.Request) {
	if err := whois.WhoisHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
