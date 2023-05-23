package services

import (
	"encoding/json"

	"github.com/Frier03/KeyAuth-API/pkg/models"
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

func (s *BadgerService) GetAllData() ([]map[string]interface{}, error) {
	var data []map[string]interface{}

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

			// Convert the APIKey struct to a map[string]interface{}
			apiKeyData := make(map[string]interface{})
			apiKeyData["id"] = apiKey.ID
			apiKeyData["subject_id"] = apiKey.SubjectID
			apiKeyData["permission_level"] = apiKey.PermissionLevel
			apiKeyData["usage"] = apiKey.Usage
			apiKeyData["limit"] = apiKey.Limit
			apiKeyData["created_at"] = apiKey.CreatedAt
			apiKeyData["expires_at"] = apiKey.ExpiresAt
			apiKeyData["last_used"] = apiKey.LastUsed
			apiKeyData["active"] = apiKey.Active

			// Append the APIKey data to the data slice
			data = append(data, apiKeyData)
		}

		return nil
	})

	return data, err
}
