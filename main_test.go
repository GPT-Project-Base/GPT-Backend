package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"context"
)


var mux = CreateAPIMux(context.Background(), "test.db")
// var srv = httptest.NewServer(mux)

func TestHelloWorldStatusCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	// Check if 
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want = %d", rec.Code, http.StatusOK)
	}
}
