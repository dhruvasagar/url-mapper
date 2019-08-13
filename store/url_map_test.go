package store_test

import (
	"testing"

	"github.com/dhruvasagar/url-mapper/store"
	"github.com/dhruvasagar/url-mapper/test"
)

func TestGetAllURLMaps(t *testing.T) {
	st := test.NewTestStore()
	defer st.Close()
	defer test.CleanTestStore()

	urlMap := store.URLMap{
		Key: "github",
		URL: "http://github.com/dhruvasagar",
	}
	st.SaveURLMap(urlMap)
	defer st.DelURLMap(urlMap.Key)

	urlMaps, err := st.GetAllURLMaps()
	if err != nil {
		t.Error("Did not expect err, got", err)
	}
	if len(urlMaps) != 1 {
		t.Error("Expected GetAllURLMaps to return 1 urlMap, got", len(urlMaps))
	}
}
