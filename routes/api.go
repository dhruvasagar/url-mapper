package routes

import (
	"github.com/dhruvasagar/url-mapper/handlers"
	"github.com/dhruvasagar/url-mapper/routes/api"
	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func InitAPI(r *mux.Router, st *store.Store) {
	s := r.PathPrefix(
		"/api/",
	).Headers(
		"Content-Type", "application/json",
	).Subrouter()

	s.Use(handlers.AuthorizationHandler)

	api.InitURLMaps(s, st)
}
