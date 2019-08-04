package store

import (
	"errors"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

func getBoltPath() string {
	dbName := os.Getenv("DB_PATH")
	if dbName == "" {
		dbName = "url-mapper.db"
	}
	return dbName
}

func getBoltBucket() string {
	dbBucket := os.Getenv("DB_BUCKET")
	if dbBucket == "" {
		dbBucket = "url_maps"
	}
	return dbBucket
}

type Store struct {
	boltDB     *bolt.DB
	boltPath   string
	boltBucket string
}

var (
	ErrNotFound = errors.New("store: key not found")
)

func Open() (*Store, error) {
	boltPath := getBoltPath()
	boltBucket := getBoltBucket()

	opts := &bolt.Options{
		Timeout: 50 * time.Millisecond,
	}
	db, err := bolt.Open(boltPath, 0640, opts)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltBucket))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &Store{
		boltDB:     db,
		boltPath:   boltPath,
		boltBucket: boltBucket,
	}, nil
}

func (s *Store) GetAll() ([]UrlMap, error) {
	values := []UrlMap{}
	s.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.boltBucket))
		b.ForEach(func(k, v []byte) error {
			values = append(values, UrlMap{
				Key: string(k),
				Url: string(v),
			})
			return nil
		})
		return nil
	})
	return values, nil
}

func (s *Store) Put(urlMap UrlMap) error {
	return s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.boltBucket))
		return b.Put([]byte(urlMap.Key), []byte(urlMap.Url))
	})
}

func (s *Store) Get(key string) (*UrlMap, error) {
	var urlMap *UrlMap
	err := s.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.boltBucket))
		val := b.Get([]byte(key))
		urlMap = &UrlMap{
			Key: key,
			Url: string(val),
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return urlMap, nil
}

func (s *Store) Delete(key string) error {
	return s.boltDB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(s.boltBucket)).Delete([]byte(key))
	})
}

func (s *Store) Close() {
	s.boltDB.Close()
}
