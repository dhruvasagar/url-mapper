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
		urlMap, err := st.GetURLMap(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if urlMap == nil {
			err := fmt.Errorf("No URL mapped to key: %q", key)
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}

		http.Redirect(w, r, urlMap.URL, http.StatusSeeOther)
	}
}

func redirectService(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		rest := vars["rest"]
		urlMap, err := st.GetURLMap(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}
		if urlMap == nil {
			err := fmt.Errorf("No URL mapped to key: %q", key)
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}
		url := fmt.Sprintf("%s/%s", urlMap.URL, rest)
		fmt.Println("URL : ", url)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}

func InitRedirect(r *mux.Router, st *store.Store) {
	r.HandleFunc("/{key}", redirect(st)).Methods("GET")
	r.HandleFunc("/{key}/{rest:.*}", redirectService(st)).Methods("GET")
}
