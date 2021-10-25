package main

import (
	"net/http"

	"github.com/yuto51942/access-tracker/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/whois", handler.WhoisHandler)
	mux.HandleFunc("/track", handler.TrackHandler)

	if err := http.ListenAndServe(":3000", mux); err != nil {
		panic(err)
	}
}
