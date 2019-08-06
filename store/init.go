package store

import (
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

type Store struct {
	db *bolt.DB
}

func New() (*Store, error) {
	boltPath := getBoltPath()
	boltURLMapsBucket := getURLMapsBucket()

	opts := &bolt.Options{
		Timeout: 50 * time.Millisecond,
	}
	db, err := bolt.Open(boltPath, 0640, opts)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(boltURLMapsBucket)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Close() {
	s.db.Close()
}
