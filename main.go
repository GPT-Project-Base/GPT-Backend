package main

import (
	"log"
	"net/http"
	"github.com/gpt-backend/routes"
	"time"
)

func NewAPIServeMux() *http.ServeMux{
	mux := http.NewServeMux()
	routes.CreateAPIEndpoints(mux)

	return mux
}

func NewAPIServer() *http.Server {
	mux := NewAPIServeMux()

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 200 * time.Millisecond,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return srv
}

func main() {
	srv := NewAPIServer()
	log.Fatal(srv.ListenAndServe())
}
