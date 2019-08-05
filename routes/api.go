package routes

import (
	"github.com/dhruvasagar/url-mapper/routes/api"
	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func InitAPI(r *mux.Router, store *store.Store) {
	s := r.PathPrefix(
		"/api/",
	).Headers(
		"Content-Type", "application/json",
	).Subrouter()

	api.InitHealthRoute(s)
	api.InitURLMapsRoutes(s, store)
}
