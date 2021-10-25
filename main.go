package main

import (
	"net/http"

	"github.com/yuto51942/access-tracker/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.TrackHandler)
	mux.HandleFunc("/whois", handler.WhoisHandler)

	if err := http.ListenAndServe(":3000", mux); err != nil {
		panic(err)
	}
}
