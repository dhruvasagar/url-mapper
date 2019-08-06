package routes

import (
	"io"
	"net/http"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "URL Mapper - Simple keyword to url mapper for short meaningful urls")
}

func InitHome(r *mux.Router, st *store.Store) {
	r.HandleFunc("/", home)
}
