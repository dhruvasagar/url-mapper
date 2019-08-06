package routes

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/gorilla/mux"
)

func setupStore() *store.Store {
	os.Setenv("DB_PATH", "url-mapper-test.db")
	st, _ := store.New()
	return st
}

func setupRedirect(dst *store.Store) (*httptest.ResponseRecorder, *store.Store) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	var st *store.Store

	if dst == nil {
		st = setupStore()
	} else {
		st = dst
	}

	InitRedirect(r, st)

	req := httptest.NewRequest("GET", "/test", nil)

	r.ServeHTTP(w, req)

	return w, st
}

func TestRedirectFailure(t *testing.T) {
	w, st := setupRedirect(nil)
	defer st.Close()

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
	st := setupStore()
	defer st.Close()

	setupURLMap(st)

	w, _ := setupRedirect(st)

	teardownURLMap(st)

	if w.Code != http.StatusSeeOther {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
	if w.Body.String() != "<a href=\"/\">See Other</a>.\n\n" {
		t.Error("Did not get expected HTTP response body, got", w.Body.String())
	}
}
