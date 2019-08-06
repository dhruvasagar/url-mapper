package routes

import (
	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, st *store.Store) {
	InitHome(r, st)
	InitHealth(r)
	InitAPI(r, st)
	InitRedirect(r, st)
}
