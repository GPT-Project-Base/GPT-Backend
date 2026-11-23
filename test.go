package main

import (
	"testing"
	"net/http/httptest"
	"context"
)


var srv = httptest.NewServer(NewAPIHandlerMux(context.Background(), "test.db"))

func TestHelloWorld(t *testing.T) {

}
