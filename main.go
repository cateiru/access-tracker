package main

import (
	"net/http"
	"os"

	"github.com/cateiru/access-tracker/routes"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var port string

func init() {
	portEnv := os.Getenv("PORT")

	if len(portEnv) == 0 {
		port = "3000"
	} else {
		port = portEnv
	}
}

func main() {
	mux := http.NewServeMux()
	h2s := &http2.Server{}

	routes.Defs(mux)

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: h2c.NewHandler(mux, h2s),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
