package main


import (
	"context"
	"time"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"fmt"
	"net/http"
	"log"
)

// API calls
func getHelloWorld(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, World!")
}



func CreateGormDatabase(dbfile string) (*gorm.DB, error) {
	// Database setup
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateAPIEndpoints(mux *http.ServeMux, db *gorm.DB) *http.ServeMux{
	// Endpoints
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		getHelloWorld(w, r, db)
	})


	return mux
}

func CreateAPIMux(ctx context.Context, dbfile string) *http.ServeMux {
	db, err := CreateGormDatabase(dbfile)
	if err != nil {
		log.Fatalf("Failed to create database: %s", err.Error())
	}

	mux := http.NewServeMux()

	CreateAPIEndpoints(mux, db)

	return mux

}

func CreateServer(ctx context.Context, dbfile string) *http.Server {
	mux := CreateAPIMux(ctx, dbfile)

	srv := &http.Server{
		Addr: 				":8080",
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
	log.Fatal(srv.ListenAndServe())
}
