package routes

import (
	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, store *store.Store) {
	InitHome(r, store)
	InitAPI(r, store)
	InitRedirect(r, store)
}
