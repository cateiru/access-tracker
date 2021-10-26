package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/yuto51942/access-tracker/handler"
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

	mux.HandleFunc("/", handler.TrackHandler)
	mux.HandleFunc("/u", handler.UserHandler)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
