package boltdb

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/yalagtyarzh/leaf_bot/repository"
)

type TokenRepository struct {
	db *bolt.DB
}

//NewTokenRepository creates a new token repository object
func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

//Save saves chatID and token in bucket
func (t *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	return t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

//Get gets the token by chatID
func (t *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}

//intToBytes returns иinary representation of a integer variable
func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
