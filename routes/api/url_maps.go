package api

import (
	"encoding/json"
	"net/http"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func renderJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func index(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlMaps, err := st.GetAllURLMaps()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderJson(w, urlMaps)
	}
}

func createOrUpdate(r *http.Request, st *store.Store) (*store.URLMap, error) {
	var urlMap store.URLMap
	err := json.NewDecoder(r.Body).Decode(&urlMap)
	if err != nil {
		return nil, err
	}
	return &urlMap, st.SaveURLMap(urlMap)
}

func create(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlMap, err := createOrUpdate(r, st)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		renderJson(w, urlMap)
	}
}

func get(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		urlMap, err := st.GetURLMap(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderJson(w, urlMap)
	}
}

func update(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlMap, err := createOrUpdate(r, st)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		renderJson(w, urlMap)
	}
}

func del(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		err := st.DelURLMap(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		renderJson(w, map[string]bool{"ok": true})
	}
}

func InitURLMaps(r *mux.Router, st *store.Store) {
	s := r.PathPrefix(
		"/url_maps/",
	).Headers(
		"Content-Type", "application/json",
	).Subrouter()

	s.HandleFunc("/", index(st)).Methods("GET")
	s.HandleFunc("/", create(st)).Methods("POST")
	s.HandleFunc("/{key}", get(st)).Methods("GET")
	s.HandleFunc("/{key}", update(st)).Methods("PUT")
	s.HandleFunc("/{key}", del(st)).Methods("DELETE")
}
