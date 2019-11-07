package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/dhruvasagar/url-mapper/test"
	"github.com/gorilla/mux"
)

func setupRedirect(dst *store.Store) (*mux.Router, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	var st *store.Store

	if dst == nil {
		st = test.NewTestStore()
	} else {
		st = dst
	}

	InitRedirect(r, st)

	return r, w
}

func setupRedirectKey(st *store.Store) (*httptest.ResponseRecorder, *store.Store) {
	r, w := setupRedirect(st)
	req := httptest.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)
	return w, st
}

func setupRedirectService(st *store.Store) (*httptest.ResponseRecorder, *store.Store) {
	r, w := setupRedirect(st)
	req := httptest.NewRequest("GET", "/test/service/route", nil)
	r.ServeHTTP(w, req)
	return w, st
}

func TestRedirectFailure(t *testing.T) {
	w, st := setupRedirectKey(nil)
	defer st.Close()
	defer test.CleanTestStore()

	if w.Code != http.StatusNotImplemented {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
	if w.Body.String() != "No URL mapped to key: \"test\"\n" {
		t.Error("Did not get expected HTTP response body, got", w.Body.String())
	}
}

func setupURLMap(st *store.Store) {
	st.SaveURLMap(store.URLMap{
		Key: "test",
		URL: "",
	})
}

func teardownURLMap(st *store.Store) {
	st.DelURLMap("test")
}

func TestRedirectSuccessful(t *testing.T) {
	st := test.NewTestStore()
	defer st.Close()
	defer test.CleanTestStore()

	setupURLMap(st)

	w, _ := setupRedirectKey(st)

	teardownURLMap(st)

	if w.Code != http.StatusSeeOther {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
	if w.Body.String() != "<a href=\"/\">See Other</a>.\n\n" {
		t.Error("Did not get expected HTTP response body, got", w.Body.String())
	}
}

func setupURLServiceMap(st *store.Store) {
	st.SaveURLMap(store.URLMap{
		Key: "test",
		URL: "http://localhost:8080",
	})
}

func TestRedirectService(t *testing.T) {
	st := test.NewTestStore()
	defer st.Close()
	defer test.CleanTestStore()

	setupURLServiceMap(st)

	w, _ := setupRedirectService(st)

	teardownURLMap(st)

	if w.Code != http.StatusSeeOther {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
	if w.Body.String() != "<a href=\"http://localhost:8080/service/route\">See Other</a>.\n\n" {
		t.Error("Did not get expected HTTP response body, got", w.Body.String())
	}
}
