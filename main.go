package main


import (
	"context"
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
	"net/http"
	"log"
)

func CreateGormDatabase(dbfile string) (*gorm.DB, error) {
	// Database setup
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddAPIEndpoints(mux *http.ServeMux, db *gorm.DB) *http.ServeMux{

	// Endpoints
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, World!")
	})

	return mux
}

func NewAPIHandlerMux(ctx context.Context, dbfile string) *http.ServeMux {
	db, err := CreateGormDatabase(dbfile)
	if err != nil {
		log.Fatal("Failed to create database.")
	}

	mux := http.NewServeMux()

	AddAPIEndpoints(mux, db)

	return mux

}

func CreateServer(ctx context.Context, dbfile string) *http.Server {

	mux := NewAPIHandlerMux(ctx, dbfile)

	srv := &http.Server{
		Handler: 			mux,
		ReadHeaderTimeout:	200 * time.Millisecond,
		ReadTimeout: 		1 * time.Second, 
		WriteTimeout: 		1 * time.Second,
		IdleTimeout:		30 * time.Second,
	}

	return srv
}

func main() {
	ctx := context.Background()

	srv := CreateServer(ctx, "main.db")
}
