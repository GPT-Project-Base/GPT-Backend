package routes

import (
	"net/http"
	"github.com/gpt-backend/api"
)

func CreateAPIEndpoints(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /hello", api.GetHelloWorld)

	return mux
}
