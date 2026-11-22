package main


import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
	"net/http"
	"log"
)

func main() {
	// Database setup
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	ctx := context.Background()



	// Test function
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
