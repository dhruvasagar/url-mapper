package store

import bolt "go.etcd.io/bbolt"

type URLMap struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func (s *Store) GetAllURLMaps() ([]URLMap, error) {
	values := []URLMap{}
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.urlMapsBucket)
		b.ForEach(func(k, v []byte) error {
			values = append(values, URLMap{
				Key: string(k),
				URL: string(v),
			})
			return nil
		})
		return nil
	})
	return values, nil
}

func (s *Store) SaveURLMap(urlMap URLMap) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.urlMapsBucket)
		return b.Put([]byte(urlMap.Key), []byte(urlMap.URL))
	})
}

func (s *Store) GetURLMap(key string) (*URLMap, error) {
	var urlMap *URLMap
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.urlMapsBucket)
		val := b.Get([]byte(key))
		urlMap = &URLMap{
			Key: key,
			URL: string(val),
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
		return tx.Bucket(s.urlMapsBucket).Delete([]byte(key))
	})
}
