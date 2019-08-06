package store

import (
	"os"

	bolt "go.etcd.io/bbolt"
)

func getURLMapsBucket() []byte {
	urlMapsBucket := os.Getenv("DB_BUCKET")
	if urlMapsBucket == "" {
		urlMapsBucket = "url_maps"
	}
	return []byte(urlMapsBucket)
}

type URLMap struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func (s *Store) GetAllURLMaps() ([]URLMap, error) {
	urlMaps := []URLMap{}
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(getURLMapsBucket())
		b.ForEach(func(k, v []byte) error {
			urlMaps = append(urlMaps, URLMap{
				Key: string(k),
				URL: string(v),
			})
			return nil
		})
		return nil
	})
	return urlMaps, nil
}

func (s *Store) SaveURLMap(urlMap URLMap) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(getURLMapsBucket())
		return b.Put([]byte(urlMap.Key), []byte(urlMap.URL))
	})
}

func (s *Store) GetURLMap(key string) (*URLMap, error) {
	var urlMap *URLMap
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(getURLMapsBucket())
		val := b.Get([]byte(key))
		if val != nil {
			urlMap = &URLMap{
				Key: key,
				URL: string(val),
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return urlMap, nil
}

func (s *Store) DelURLMap(key string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(getURLMapsBucket()).Delete([]byte(key))
	})
}
