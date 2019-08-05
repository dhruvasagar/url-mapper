package routes

import (
	"fmt"
	"net/http"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func redirect(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		urlMap, err := st.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if urlMap == nil {
			err := fmt.Errorf("No URL mapped to key: %s", key)
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}

		http.Redirect(w, r, urlMap.Url, http.StatusSeeOther)
	}
}

func InitRedirect(r *mux.Router, store *store.Store) {
	r.HandleFunc("/{key}", redirect(store)).Methods("GET")
}
