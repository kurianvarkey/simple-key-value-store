package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type fileStore struct {
	store map[string]string
}

// NewStore returns a new Store
func NewFileStore() StoreInterface {
	return &fileStore{
		store: make(map[string]string),
	}
}

// Set a key-value pair
func (s *fileStore) Set(key string, value string) error {
	s.store[key] = value
	return nil
}

// Get the value for a key
func (s *fileStore) Get(key string) (string, error) {
	if value, ok := s.store[key]; ok {
		return value, nil
	}

	return "", fmt.Errorf("key %s not found", key)
}

// Delete a key-value pair
func (s *fileStore) Delete(key string) error {
	if _, ok := s.store[key]; ok {
		delete(s.store, key)
		return nil
	}

	return fmt.Errorf("key %s not found", key)
}

// List all key-value pairs
func (s *fileStore) List() (map[string]string, error) {
	if len(s.store) == 0 {
		return nil, fmt.Errorf("store is empty")
	}

	return s.store, nil
}

// save the store to a file
func (s *fileStore) Save() error {
	file := "store.json"

	data, err := json.MarshalIndent(s.store, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}

// load the store from a file
func (s *fileStore) Load() error {
	file := "store.json"

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.store)
}
