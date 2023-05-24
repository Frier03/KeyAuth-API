package services

import (
	"encoding/json"
	"time"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/Frier03/KeyAuth-API/pkg/utils"
	"github.com/dgraph-io/badger/v3"
)

type BadgerService struct {
	db *badger.DB
}

func NewBadgerService() (*BadgerService, error) {
	// Specify the options for the database
	options := badger.DefaultOptions("").WithInMemory(true) // in-memory database

	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	return &BadgerService{
		db: db,
	}, nil
}

func (s *BadgerService) Close() error {
	return s.db.Close()
}

func (s *BadgerService) PutAPIKey(key []byte, apiKey *models.APIKey) error {
	// Serialize the APIKey struct into a byte slice
	value, err := json.Marshal(apiKey)
	if err != nil {
		return err
	}

	err = s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})

	return err
}

func (s *BadgerService) Get(key []byte) ([]byte, error) {
	var result []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		value, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		result = value
		return nil
	})

	return result, err
}

func (s *BadgerService) GetAllData() ([]models.APIKey, error) {
	var apiKeys []models.APIKey

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()

			// Retrieve the value for the current key
			value, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			// Deserialize the value into an APIKey struct
			var apiKey models.APIKey
			err = json.Unmarshal(value, &apiKey)
			if err != nil {
				return err
			}

			// Apply necessary formatting to the APIKey struct
			apiKey.CreatedAt = utils.FormatCreatedAt(apiKey.CreatedAt)
			apiKey.ExpiresAt = utils.FormatExpiresAt(apiKey.ExpiresAt)
			apiKey.LastUsed = utils.FormatUsedAt(apiKey.LastUsed)

			// Append the APIKey to the apiKeys slice
			apiKeys = append(apiKeys, apiKey)
		}

		return nil
	})

	return apiKeys, err
}

func (s *BadgerService) FetchTotalGeneratedAPIKeys() int {
	count := 0

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10 // Adjust the prefetch size for performance

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			count++
		}

		return nil
	})

	if err != nil {
		return 0
	}

	return count
}

func (s *BadgerService) FetchTotalExpiredAPIKeys() int {
	count := 0
	timeLayout := "2006-01-02 15:04:05.999999999 -0700 MST"

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10 // Adjust the prefetch size for performance

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()

			// Retrieve the ExpiresAt value from the key's metadata
			expiresAt := item.UserMeta()

			// Parse the ExpiresAt value as a time.Time
			t, err := time.Parse(timeLayout, string(expiresAt))
			if err != nil {
				return err
			}

			// Check if the key has expired
			if t.Before(time.Now()) {
				count++
			}
		}

		return nil
	})

	if err != nil {
		return 0
	}

	return count
}
