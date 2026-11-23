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

func CreateAPIEndpoints(mux *http.ServeMux, db *gorm.DB) {

	// Endpoints
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}

func CreateServer(ctx context.Context) *http.Server {

	db, err := CreateGormDatabase("main.db")
	mux := http.NewServeMux()

	CreateAPIEndpoints(mux, db)

	if err != nil {
		log.Fatal("Failed to create database.")
	}

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

	srv := CreateServer(ctx)

}
