package web

import (
	"golang-hatena-client/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Top).Methods("GET")
	r.HandleFunc("/rss", handler.GetRSS).Methods("GET")

	return r
}
