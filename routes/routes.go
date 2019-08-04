package routes

import (
	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, store *store.Store) {
	InitHomeRoutes(r, store)
	InitAPIRoutes(r, store)
	InitRedirectRoutes(r, store)
}
