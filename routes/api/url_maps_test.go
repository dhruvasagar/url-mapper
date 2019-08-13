package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/dhruvasagar/url-mapper/test"
	"github.com/gorilla/mux"
)

func setupURLMapsAPI() (*httptest.ResponseRecorder, *mux.Router, *store.Store) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	st := test.NewTestStore()
	st.SaveURLMap(store.URLMap{
		Key: "test",
		URL: "",
	})

	InitURLMaps(r, st)
	return w, r, st
}

func TestIndex(t *testing.T) {
	w, r, st := setupURLMapsAPI()
	defer st.Close()
	defer test.CleanTestStore()

	req := httptest.NewRequest("GET", "/url_maps/", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	header := w.Header().Get("Content-Type")
	if header != "application/json" {
		t.Error("Did not get expected HTTP header Content-Type, got", header)
	}

	var urlMaps []store.URLMap
	json.NewDecoder(w.Body).Decode(&urlMaps)
	if len(urlMaps) != 1 {
		t.Error("Did not get expected response body, got", w.Body.String())
	}
}

func TestCreate(t *testing.T) {
	w, r, st := setupURLMapsAPI()
	defer st.Close()
	defer test.CleanTestStore()

	body, _ := json.Marshal(store.URLMap{
		Key: "github",
		URL: "https://github.com/dhruvasagar",
	})

	req := httptest.NewRequest("POST", "/url_maps/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
}

func TestGet(t *testing.T) {
	w, r, st := setupURLMapsAPI()
	defer st.Close()
	defer test.CleanTestStore()

	req := httptest.NewRequest("GET", "/url_maps/test", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	var urlMap store.URLMap
	body := w.Body.String()
	json.NewDecoder(w.Body).Decode(&urlMap)
	if urlMap.Key != "test" || urlMap.URL != "" {
		t.Error("Did not get expected HTTP response body, got", body)
	}
}

func TestUpdate(t *testing.T) {
	w, r, st := setupURLMapsAPI()
	defer st.Close()
	defer test.CleanTestStore()

	body, _ := json.Marshal(store.URLMap{
		Key: "test",
		URL: "https://dhruvasagar.dev",
	})

	req := httptest.NewRequest("PUT", "/url_maps/test", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	var rURLMap store.URLMap
	rbody := w.Body.String()
	json.NewDecoder(w.Body).Decode(&rURLMap)
	if rURLMap.Key != "test" || rURLMap.URL != "https://dhruvasagar.dev" {
		t.Error("Did not get expected HTTP response body, got", rbody)
	}

	dURLMap, _ := st.GetURLMap("test")
	if dURLMap.URL != "https://dhruvasagar.dev" {
		t.Error("Did not update urlMap")
	}
}

func TestDelete(t *testing.T) {
	w, r, st := setupURLMapsAPI()
	defer st.Close()
	defer test.CleanTestStore()

	req := httptest.NewRequest("DELETE", "/url_maps/test", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	if w.Body.String() != "{\"ok\":true}\n" {
		t.Error("Did not get expected HTTP response body, got", w.Body.String())
	}

	urlMap, _ := st.GetURLMap("test")
	if urlMap != nil {
		t.Error("Did not delete urlMap")
	}
}
