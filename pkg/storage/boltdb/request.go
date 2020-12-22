package boltdb

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/storage"
	"strconv"
)

type TokenStorage struct {
	db     *bolt.DB
}

func NewTokenStorage(db *bolt.DB) *TokenStorage {
	return &TokenStorage{db: db}
}

func (s *TokenStorage) Save(chatID int64, token string, bucket storage.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (s *TokenStorage) Get(chatID int64, bucket storage.Bucket) (string, error) {
	var token string

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		token = string(b.Get(intToBytes(chatID)))
		return nil
	})

	if token == "" {
		return "", errors.New("not found")
	}

	return token, err
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
