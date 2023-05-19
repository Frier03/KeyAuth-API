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
