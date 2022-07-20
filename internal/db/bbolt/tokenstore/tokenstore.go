package tokenstore

import (
	"errors"
	"github.com/audetv/telegram-bot-getpocket/internal/pkg/repos/token"
	"go.etcd.io/bbolt"
	"strconv"
)

type TokenRepository struct {
	db *bbolt.DB
}

func NewTokenRepository(db *bbolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Create(chatID int64, token string, bucket token.Bucket) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (r *TokenRepository) Get(chatID int64, bucket token.Bucket) (string, error) {
	var tokenValue string

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		tokenValue = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	if tokenValue == "" {
		return "", errors.New("token not found")
	}

	return tokenValue, nil
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
