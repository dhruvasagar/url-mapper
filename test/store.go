package test

import (
	"os"

	"github.com/dhruvasagar/url-mapper/store"
)

func getDBPath() string {
	return "url-mapper-test.db"
}

func NewTestStore() *store.Store {
	os.Setenv("DB_PATH", getDBPath())
	st, _ := store.New()
	return st
}

func CleanTestStore() {
	os.Remove(getDBPath())
}
