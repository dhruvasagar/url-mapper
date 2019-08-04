package routes

import (
	"net/http"
	"os"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func getDefaultRoute() string {
	defaultRoute := os.Getenv("DEFAULT_ROUTE")
	if defaultRoute == "" {
		defaultRoute = "https://dhruvasagar.dev"
	}
	return defaultRoute
}

func InitRedirectRoutes(r *mux.Router, store *store.Store) {
	r.HandleFunc("/{key}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		if key == "" {
			http.Redirect(w, r, getDefaultRoute(), http.StatusSeeOther)
			return
		}

		urlMap, err := store.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if urlMap == nil {
			http.Redirect(w, r, getDefaultRoute(), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, urlMap.Url, http.StatusSeeOther)
	})
}
