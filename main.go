package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/yuto51942/access-tracker/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var port string

func init() {
	_port := os.Getenv("PORT")

	if len(_port) == 0 {
		port = ":3000"
	} else {
		port = strings.Join([]string{":", _port}, "")
	}
}

func main() {
	mux := http.NewServeMux()
	h2s := &http2.Server{}

	mux.HandleFunc("/", handler.TrackHandler)
	mux.HandleFunc("/u", handler.UserHandler)

	server := &http.Server{
		Addr:    strings.Join([]string{"0.0.0.0", port}, ""),
		Handler: h2c.NewHandler(mux, h2s),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
