package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	st := &store.Store{}

	r := mux.NewRouter()
	InitHome(r, st)

	req := httptest.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
	if w.Body.String() != "URL Mapper - Simple keyword to url mapper for short meaningful urls" {
		t.Error("Did not get expected HTTP response body, got", w.Body.String())
	}
}
