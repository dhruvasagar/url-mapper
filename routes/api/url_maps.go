package api

import (
	"encoding/json"
	"net/http"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func index(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlMaps, err := st.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(urlMaps)
	}
}

func createOrUpdate(r *http.Request, st *store.Store) (*store.UrlMap, error) {
	var urlMap store.UrlMap
	err := json.NewDecoder(r.Body).Decode(&urlMap)
	if err != nil {
		return nil, err
	}
	return &urlMap, st.Put(urlMap)
}

func create(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlMap, err := createOrUpdate(r, st)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(urlMap)
	}
}

func get(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		urlMap, err := st.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(urlMap)
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
		json.NewEncoder(w).Encode(urlMap)
	}
}

func del(st *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		err := st.Delete(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func InitURLMapsRoutes(r *mux.Router, store *store.Store) {
	s := r.PathPrefix(
		"/url_maps/",
	).Headers(
		"Content-Type", "application/json",
	).Subrouter()

	s.HandleFunc("/", index(store)).Methods("GET")
	s.HandleFunc("/", create(store)).Methods("POST")
	s.HandleFunc("/{key}", get(store)).Methods("GET")
	s.HandleFunc("/{key}", update(store)).Methods("PUT")
	s.HandleFunc("/{key}", del(store)).Methods("DELETE")
}
